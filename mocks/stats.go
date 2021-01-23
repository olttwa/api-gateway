package mocks

import "github.com/stretchr/testify/mock"

type Stats struct {
	mock.Mock
}

func (s *Stats) Success() int {
	args := s.Called()
	return args.Int(0)
}

func (s *Stats) Errors() int {
	args := s.Called()
	return args.Int(0)
}

func (s *Stats) Mean() int {
	args := s.Called()
	return args.Int(0)
}

func (s *Stats) Percentile(p float64) int {
	args := s.Called(p)
	return args.Int(0)
}
