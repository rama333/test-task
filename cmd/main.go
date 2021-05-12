package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"test-task/services"
	"time"
)

func main() {

	if err := run(); err != nil {
		logrus.Fatal("failed run")
	}

}

type ServiceRepo struct {
	service services.Service
}

func run() error {

	var st time.Time
	var n uint64
	var p time.Duration

	defer func() {
		logrus.WithField("shutdown_time", time.Now().Sub(st)).Info("stopped")
	}()

	ctx := context.Background()

	items := make(chan services.Item, n)

	go func(item chan services.Item) {
		for {
			item <- struct{}{}
			time.Sleep(100 * time.Millisecond)
		}
	}(items)

	ser := ServiceRepo{}

	n, p = ser.service.GetLimits()

	go func(item chan services.Item) {
		for {
			select {
			case <-ctx.Done():
				break
			default:
			}

			bath := Batches(items, n, p)
			ser.service.Process(ctx, bath)
		}
	}(items)

	defer close(items)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	return nil
}

func Batches(items chan services.Item, maxItems uint64, maxTimes time.Duration) services.Batch {

	batches := make(services.Batch, 0)

LOOP:
	for {
		tim := time.After(maxTimes)
		select {
		case val, ok := <-items:

			if !ok {
				break LOOP
			}
			batches = append(batches, val)

			if uint64(len(batches)) == maxItems {
				break LOOP
			}

		case <-tim:
			break LOOP
		}
	}

	return batches
}
