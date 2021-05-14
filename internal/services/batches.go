package services

import (
	"context"
	"test-task/internal/entity"
	"time"
)

type Batch struct {
	service ObjectService
	maxItems uint64
	maxTimes time.Duration

}

func NewBatch(service ObjectService) (*Batch)  {

	n, p := service.GetLimits()

	return &Batch{service:service, maxItems: n, maxTimes: p}
}

func (b *Batch) RunBatches(items chan entity.Item, ctx context.Context)  {
	go func(item chan entity.Item) {
		for {
			select {
			case <-ctx.Done():
				break
			default:
			}

			bath := b.batches(items)
			b.service.Process(ctx, bath)
		}
	}(items)
}

func (b *Batch) batches(items chan entity.Item) entity.Batch {

	batches := make(entity.Batch, 0)

LOOP:
	for {
		tim := time.After(b.maxTimes)
		select {
		case val, ok := <-items:

			if !ok {
				break LOOP
			}
			batches = append(batches, val)

			if uint64(len(batches)) == b.maxItems {
				break LOOP
			}

		case <-tim:
			break LOOP
		}
	}

	return batches
}