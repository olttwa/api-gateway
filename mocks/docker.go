package mocks

import (
	"context"
	"rgate/model"

	"github.com/stretchr/testify/mock"
)

type Docker struct {
	mock.Mock
}

func (d *Docker) ListContainers(ctx context.Context, ml []string) ([]model.Container, error) {
	args := d.Called(ctx, ml)
	return args.Get(0).([]model.Container), args.Error(1)
}
