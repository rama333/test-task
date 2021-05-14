package services

import (
	"context"
	"test-task/configs"
	"test-task/internal/entity"
	"time"
)

type ServiceProcces struct {
	n uint64
	p time.Duration
}

func NewServiceProcces(conf configs.Config) (*ServiceProcces)  {

	return &(ServiceProcces{n:conf.N, p: conf.P})
}

func (s *ServiceProcces) GetLimits() (n uint64, p time.Duration) {

	return s.n, s.p
}
func (s *ServiceProcces) Process(ctx context.Context, batch entity.Batch) error {

	time.Sleep(100 * time.Millisecond)

	return nil
}