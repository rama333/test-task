package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"test-task/configs"
	"test-task/internal/entity"
	"test-task/internal/services"
	"time"
)

func main() {

	if err := run(); err != nil {
		logrus.Fatal(err)
	}

}

func run() error {

	var st time.Time

	defer func() {
		logrus.WithField("shutdown_time", time.Now().Sub(st)).Info("stopped")
	}()

	config, err := configs.LoadConfig("../configs/conf.conf")

	if err != nil {
		return errors.Wrap(err, "failed load config")
	}

	err = config.Validate()
	if err != nil {
		return errors.Wrap(err, "failed validate config")
	}

	process := services.NewServiceProcces(config)

	ctx := context.Background()

	items := make(chan entity.Item, config.N)

	defer close(items)

	go func(item chan entity.Item) {
		for {
			item <- struct{}{}
			time.Sleep(100 * time.Millisecond)
		}
	}(items)

	batch := services.NewBatch(process)

	batch.RunBatches(items, ctx)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	return nil
}


