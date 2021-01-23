package utils_test

import (
	"rgate/model"
	"rgate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchRoutesToBackend(t *testing.T) {
	s := model.Spec{
		Routes: []model.RouteSpec{
			{
				Backend:    "payment",
				PathPrefix: "/api/payment",
			},
			{
				Backend:    "orders",
				PathPrefix: "/api/orders",
			},
		},
		Backends: []model.Backend{
			{
				Name:        "payment",
				MatchLabels: []string{"app=payment", "version=v2"},
			},
			{
				Name:        "orders",
				MatchLabels: []string{"app=orders"},
			},
		},
	}

	r, err := utils.MatchRoutesToBackend(s)
	assert.NoError(t, err)

	expected := []model.Route{
		{
			PathPrefix:  "/api/payment",
			MatchLabels: []string{"app=payment", "version=v2"},
		},
		{
			PathPrefix:  "/api/orders",
			MatchLabels: []string{"app=orders"},
		},
	}
	assert.Equal(t, expected, r)
}

func TestMatchRoutesToBackendError(t *testing.T) {
	s := model.Spec{
		Routes: []model.RouteSpec{
			{
				Backend:    "missing",
				PathPrefix: "/api/missing",
			},
		},
	}

	_, err := utils.MatchRoutesToBackend(s)
	assert.EqualError(t, err, "backend not defined for route: /api/missing")
}
