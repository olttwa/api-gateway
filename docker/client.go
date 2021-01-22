package docker

import (
	"context"
	"fmt"
	"log"
	"rgate/model"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Docker struct {
	c *client.Client
}

func (d Docker) ListContainers(ctx context.Context, labels []string) ([]model.Container, error) {
	filters := filters.NewArgs()
	for _, l := range labels {
		filters.Add("label", l)
	}

	containers, err := d.c.ContainerList(ctx, types.ContainerListOptions{Filters: filters})
	if err != nil {
		return nil, err
	}

	c := make([]model.Container, 0, len(containers))
	for _, container := range containers {
		for _, p := range container.Ports {
			c = append(c, model.Container{
				IP:   p.IP,
				Port: fmt.Sprintf("%d", p.PublicPort),
			})
		}

	}
	return c, nil
}

func New() Docker {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error initializing docker client: %s", err)
	}
	return Docker{c}
}
