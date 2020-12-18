package core

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/opctl/opctl/sdks/go/internal/uniquestring"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/pubsub"
)

//counterfeiter:generate -o internal/fakes/parallelCaller.go . parallelCaller
type parallelCaller interface {
	// Executes a parallel call
	Call(
		parentCtx context.Context,
		callID string,
		inboundScope map[string]*model.Value,
		rootCallID string,
		opPath string,
		callSpecParallelCall []*model.CallSpec,
	) (
		map[string]*model.Value,
		error,
	)
}

func newParallelCaller(
	caller caller,
	pubSub pubsub.PubSub,
) parallelCaller {

	return _parallelCaller{
		caller:              caller,
		pubSub:              pubSub,
		uniqueStringFactory: uniquestring.NewUniqueStringFactory(),
	}

}

func refToName(ref string) string {
	return strings.TrimSuffix(strings.TrimPrefix(ref, "$("), ")")
}

type _parallelCaller struct {
	caller              caller
	pubSub              pubsub.PubSub
	uniqueStringFactory uniquestring.UniqueStringFactory
}

func (pc _parallelCaller) Call(
	parentCtx context.Context,
	callID string,
	inboundScope map[string]*model.Value,
	rootCallID string,
	opPath string,
	callSpecParallelCall []*model.CallSpec,
) (
	map[string]*model.Value,
	error,
) {
	// setup cancellation
	parallelCtx, cancelParallel := context.WithCancel(parentCtx)
	defer cancelParallel()

	childCallNeededCountByName := map[string]int{}
	for _, callSpecChildCall := range callSpecParallelCall {
		// increment needed by counts for any needs
		for _, neededCallRef := range callSpecChildCall.Needs {
			childCallNeededCountByName[refToName(neededCallRef)]++
		}
	}

	startTime := time.Now().UTC()
	childCallIndexByID := map[string]int{}
	childCallCancellationByName := map[string]context.CancelFunc{}
	childCallOutputsByIndex := map[int]map[string]*model.Value{}

	// perform calls in parallel w/ cancellation
	for childCallIndex, childCall := range callSpecParallelCall {
		childCtx := parallelCtx

		childCallID, err := pc.uniqueStringFactory.Construct()
		if nil != err {
			// end run immediately on any error
			return nil, err
		}

		childCallIndexByID[childCallID] = childCallIndex

		if nil != childCall.Name {
			newCtx, cancel := context.WithCancel(childCtx)
			childCtx = newCtx

			childCallCancellationByName[*childCall.Name] = cancel
		}

		go func(childCall *model.CallSpec) {
			defer func() {
				if panicArg := recover(); panicArg != nil {
					// recover from panics; treat as errors
					fmt.Printf("%v\n%v", panicArg, debug.Stack())

					// cancel all children on any error
					cancelParallel()
				}
			}()

			pc.caller.Call(
				childCtx,
				childCallID,
				inboundScope,
				childCall,
				opPath,
				&callID,
				rootCallID,
			)

		}(childCall)
	}

	// subscribe to events
	// @TODO: handle err channel
	eventChannel, _ := pc.pubSub.Subscribe(
		// don't cancel w/ children; we need to read err msgs
		parentCtx,
		model.EventFilter{
			Roots: []string{rootCallID},
			Since: &startTime,
		},
	)

	var isChildErred = false
	outputs := map[string]*model.Value{}

eventLoop:
	for event := range eventChannel {
		if nil != event.CallEnded {
			if childCallIndex, isChildCallEnded := childCallIndexByID[event.CallEnded.Call.ID]; isChildCallEnded {
				childCallOutputsByIndex[childCallIndex] = event.CallEnded.Outputs
				if nil != event.CallEnded.Error {
					isChildErred = true

					// cancel all children on any error
					cancelParallel()
				}

				// decrement needed by counts for any needs
				for _, neededCallRef := range callSpecParallelCall[childCallIndex].Needs {
					childCallNeededCountByName[refToName(neededCallRef)]--
				}

				for neededCallName, neededCount := range childCallNeededCountByName {
					if 1 > neededCount {
						if cancel, ok := childCallCancellationByName[neededCallName]; ok {
							cancel()
						}
					}
				}
			}

			if len(childCallOutputsByIndex) == len(childCallIndexByID) {
				// all calls have ended

				// construct parallel outputs
				for i := 0; i < len(callSpecParallelCall); i++ {
					callOutputs := childCallOutputsByIndex[i]
					for varName, varData := range callOutputs {
						outputs[varName] = varData
					}
				}

				if isChildErred {
					return nil, errors.New("child call failed")
				}

				break eventLoop
			}

		}
	}

	return outputs, nil
}
