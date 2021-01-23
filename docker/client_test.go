package docker_test

import (
	"context"
	"io"
	"os"
	"rgate/docker"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"gotest.tools/v3/assert"
)

func TestListContainers(t *testing.T) {
	ctx := context.Background()
	// Start a container with labels.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	assert.NilError(t, err)
	imageName := "nginx"

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	assert.NilError(t, err)
	_, err = io.Copy(os.Stdout, out)
	assert.NilError(t, err)

	ports, bindings, err := nat.ParsePortSpecs([]string{"5001:80"})
	assert.NilError(t, err)

	hostConfig := &container.HostConfig{
		PortBindings: bindings,
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,
		ExposedPorts: ports,
		Labels:       map[string]string{"rgate": "test"},
	}, hostConfig, nil, nil, "")
	assert.NilError(t, err)

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	assert.NilError(t, err)

	// Remove running containers.
	defer func() {
		err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})
		assert.NilError(t, err)
	}()

	// List containers having given labels.
	d := docker.New()
	c, err := d.ListContainers(ctx, []string{"rgate=test"})
	assert.NilError(t, err)
	assert.Equal(t, "0.0.0.0", c[0].IP)
	assert.Equal(t, "5001", c[0].Port)
}
