package services

import (
	"context"
	"test-task/internal/entity"
	"time"
)

// Service defines external service that can process batches of items.
type ObjectService interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch entity.Batch) error
}



