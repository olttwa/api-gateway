package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type docker struct {
	c *client.Client
}

type Client interface {
	ListContainers(ctx context.Context, labels []string) ([]Container, error)
}

func Initialize() Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error initializing docker client: %s", err)
	}
	return docker{cli}
}

func (d docker) ListContainers(ctx context.Context, labels []string) ([]Container, error) {
	filters := filters.NewArgs()
	for _, l := range labels {
		filters.Add("label", l)
	}

	containers, err := d.c.ContainerList(ctx, types.ContainerListOptions{Filters: filters})
	if err != nil {
		return nil, err
	}

	c := make([]Container, 0, len(containers))
	for _, container := range containers {
		for _, p := range container.Ports {
			c = append(c, Container{
				IP:   p.IP,
				Port: fmt.Sprintf("%d", p.PublicPort),
			})
		}

	}
	return c, nil
}
