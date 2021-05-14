package objectservice

import (
	"test-task/internal/entity"
	"time"
)

type Batch struct {
	maxItems uint64
	maxTimes time.Duration
}

func NewBatch(maxItems uint64, maxTimes time.Duration) (*Batch)  {

	return &Batch{maxItems: maxItems, maxTimes: maxTimes}
}

func (b *Batch) Batches(items chan entity.Item) entity.Batch {

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