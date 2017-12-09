package docker

//go:generate counterfeiter -o ./fakeImagePuller.go --fake-name fakeImagePuller ./ imagePuller

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	dockerClientPkg "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/util/pubsub"
	"golang.org/x/net/context"
	"io"
)

type imagePuller interface {
	Pull(
		dcgContainerImage *model.DCGContainerCallImage,
		containerId string,
		rootOpId string,
		eventPublisher pubsub.EventPublisher,
	) error
}

func newImagePuller(
	dockerClient dockerClientPkg.CommonAPIClient,
) imagePuller {
	return _imagePuller{
		dockerClient,
	}
}

type _imagePuller struct {
	dockerClient dockerClientPkg.CommonAPIClient
}

func (ip _imagePuller) Pull(
	dcgContainerImage *model.DCGContainerCallImage,
	containerId string,
	rootOpId string,
	eventPublisher pubsub.EventPublisher,
) error {

	imagePullOptions := types.ImagePullOptions{}
	if nil != dcgContainerImage.PullCreds &&
		"" != dcgContainerImage.PullCreds.Username &&
		"" != dcgContainerImage.PullCreds.Password {
		var err error
		imagePullOptions.RegistryAuth, err = constructRegistryAuth(
			dcgContainerImage.PullCreds.Username,
			dcgContainerImage.PullCreds.Password,
		)
		if nil != err {
			return err
		}
	}

	imagePullResp, err := ip.dockerClient.ImagePull(
		context.Background(),
		dcgContainerImage.Ref,
		imagePullOptions,
	)
	if nil != err {
		return err
	}
	defer imagePullResp.Close()

	stdOutWriter := NewStdOutWriteCloser(eventPublisher, containerId, rootOpId)
	defer stdOutWriter.Close()

	dec := json.NewDecoder(imagePullResp)
	for {
		var jm jsonmessage.JSONMessage
		if err = dec.Decode(&jm); nil != err {
			if err == io.EOF {
				err = nil
			}
			return err
		}
		jm.Display(stdOutWriter, nil)
	}
}