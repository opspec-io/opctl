package docker

//go:generate counterfeiter -o ./fakeDockerClient.go --fake-name fakeDockerClient ./ dockerClient

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"golang.org/x/net/context"
	"io"
)

// client interface for docker
type dockerClient interface {
	// ContainerCreate creates a new container based in the given configuration.
	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string,
	) (container.ContainerCreateCreatedBody, error)

	// ContainerInspect returns the container information.
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)

	// ContainerLogs returns the logs generated by a container in an io.ReadCloser.
	ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error)

	// ContainerRemove kills and removes a container from the docker host.
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error

	// ContainerStart sends a request to the docker daemon to start a container.
	ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error

	// ContainerWait pauses execution until a container exits.
	ContainerWait(ctx context.Context, containerID string) (int64, error)

	// ImagePull requests the docker host to pull an image from a remote registry.
	// It executes the privileged function if the operation is unauthorized
	// and it tries one more time.
	ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error)
}

type dockerNotFoundError struct {
	error
}

func (this dockerNotFoundError) NotFound() bool {
	return true
}
