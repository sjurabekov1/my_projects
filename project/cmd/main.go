package main

import (
	"context"

	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	grpcMain "google.golang.org/grpc"

	"hyssa/hyssa_go_billing_service/internal/config"
	"hyssa/hyssa_go_billing_service/internal/core/repository"
	"hyssa/hyssa_go_billing_service/internal/pkg/logger"
	"hyssa/hyssa_go_billing_service/internal/transport/grpc"
	"hyssa/hyssa_go_billing_service/pkg/logger/factory"
)

var (
	cfg *config.Config
)

func init() {
	log.Println("Initializing...")
	cfg = config.Load()
	loggerFactory, err := factory.Build(&cfg.Logging)
	if err != nil {
		panic(err)
	}
	logger.SetLogger(loggerFactory)
	logger.Log.Info("Initializing done...")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Project.Timout)
	defer cancel()

	repos := repository.New(ctx, cfg)
	grpcServer := grpc.New(repos)

	lis, err := net.Listen("tcp", cfg.GRPC.URL)
	if err != nil {
		log.Fatal("Error while listening: ", err)
		return
	}

	go func() {
		logger.Log.Info("starting grpc server on " + cfg.GRPC.URL)
		grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	gracefulShutdown(grpcServer, ctx, cancel)
}

func gracefulShutdown(grpcServer *grpcMain.Server, ctx context.Context, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		logger.Log.Info("shutting down")

		grpcServer.GracefulStop()

		logger.Log.Info("shutdown successfully called")
		wg.Done()
	}(&wg)

	go func() {
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
	os.Exit(0)
}
