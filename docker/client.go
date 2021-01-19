package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type docker struct {
	c *client.Client
}

type Client interface {
	ListContainers(ctx context.Context, labels []string) ([]types.Container, error)
}

func Initialize() Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error initializing docker client: %s", err)
	}
	return docker{cli}
}

func (d docker) ListContainers(ctx context.Context, labels []string) ([]types.Container, error) {
	filters := filters.NewArgs()
	for _, l := range labels {
		filters.Add("label", l)
	}
	containers, err := d.c.ContainerList(ctx, types.ContainerListOptions{Filters: filters})
	if err != nil {
		return nil, err
	}
	return containers, nil
}
