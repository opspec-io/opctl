package container

import (
	"path/filepath"

	"github.com/golang-interfaces/ios"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container/cmd"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container/dirs"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container/envvars"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container/files"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container/image"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container/sockets"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/str"
)

//counterfeiter:generate -o fakes/interpreter.go . Interpreter
type Interpreter interface {
	// Interpret interprets an CallContainerSpec into a DCGContainerCall
	Interpret(
		scope map[string]*model.Value,
		callContainerSpec *model.CallContainerSpec,
		containerID string,
		rootOpID string,
		opPath string,
	) (*model.DCGContainerCall, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter(
	dataDirPath string,
) Interpreter {
	return _interpreter{
		cmdInterpreter:     cmd.NewInterpreter(),
		dirsInterpreter:    dirs.NewInterpreter(dataDirPath),
		envVarsInterpreter: envvars.NewInterpreter(),
		filesInterpreter:   files.NewInterpreter(dataDirPath),
		imageInterpreter:   image.NewInterpreter(),
		os:                 ios.New(),
		dataDirPath:        dataDirPath,
		stringInterpreter:  str.NewInterpreter(),
		socketsInterpreter: sockets.NewInterpreter(),
	}
}

type _interpreter struct {
	cmdInterpreter     cmd.Interpreter
	dirsInterpreter    dirs.Interpreter
	envVarsInterpreter envvars.Interpreter
	filesInterpreter   files.Interpreter
	imageInterpreter   image.Interpreter
	os                 ios.IOS
	dataDirPath        string
	stringInterpreter  str.Interpreter
	socketsInterpreter sockets.Interpreter
}

func (itp _interpreter) Interpret(
	scope map[string]*model.Value,
	callContainerSpec *model.CallContainerSpec,
	containerID string,
	rootOpID string,
	opPath string,
) (*model.DCGContainerCall, error) {

	dcgContainerCall := &model.DCGContainerCall{
		DCGBaseCall: model.DCGBaseCall{
			RootOpID: rootOpID,
			OpPath:   opPath,
		},
		Dirs:        map[string]string{},
		EnvVars:     map[string]string{},
		Files:       map[string]string{},
		Sockets:     map[string]string{},
		WorkDir:     callContainerSpec.WorkDir,
		ContainerID: containerID,
		Ports:       callContainerSpec.Ports,
	}

	// construct dcg container path
	scratchDirPath := filepath.Join(
		itp.dataDirPath,
		"dcg",
		rootOpID,
		"containers",
		containerID,
		"fs",
	)
	if err := itp.os.MkdirAll(scratchDirPath, 0700); nil != err {
		return nil, err
	}

	// interpret cmd
	var err error
	dcgContainerCall.Cmd, err = itp.cmdInterpreter.Interpret(scope, callContainerSpec.Cmd)
	if nil != err {
		return nil, err
	}

	// interpret dirs
	dcgContainerCall.Dirs, err = itp.dirsInterpreter.Interpret(scope, callContainerSpec.Dirs, scratchDirPath)
	if nil != err {
		return nil, err
	}

	// interpret envVars
	dcgContainerCall.EnvVars, err = itp.envVarsInterpreter.Interpret(scope, callContainerSpec.EnvVars)
	if nil != err {
		return nil, err
	}

	// interpret files
	dcgContainerCall.Files, err = itp.filesInterpreter.Interpret(scope, callContainerSpec.Files, scratchDirPath)
	if nil != err {
		return nil, err
	}

	// interpret image
	dcgContainerCall.Image, err = itp.imageInterpreter.Interpret(scope, callContainerSpec.Image, scratchDirPath)
	if nil != err {
		return nil, err
	}

	// interpret name as string
	if nil != callContainerSpec.Name {
		dcgContainerCallName, err := itp.stringInterpreter.Interpret(scope, *callContainerSpec.Name)
		if nil != err {
			return nil, err
		}
		dcgContainerCall.Name = dcgContainerCallName.String
	}

	// interpret workDir
	if "" != callContainerSpec.WorkDir {
		dcgContainerCallWorkDir, err := itp.stringInterpreter.Interpret(scope, callContainerSpec.WorkDir)
		if nil != err {
			return nil, err
		}

		dcgContainerCall.WorkDir = *dcgContainerCallWorkDir.String
	}

	// interpret sockets
	dcgContainerCall.Sockets, err = itp.socketsInterpreter.Interpret(scope, callContainerSpec.Sockets, scratchDirPath)

	return dcgContainerCall, err

}
