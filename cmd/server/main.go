package main

import (
	"context"
	"fibonachi/internal/config"
	"fibonachi/internal/delivery/http/handler"
	"fibonachi/internal/server/rest"
	"fibonachi/internal/service"
	"fibonachi/internal/store"
	"fibonachi/internal/store/redisclient"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := config.New()

	client, err := redisclient.GetClient(cfg.RedisAddress)
	if err != nil {
		logrus.Fatal(err)
	}

	stores := store.NewStore(client)
	services := service.NewService(stores)
	handlers := handler.New(services)

	srv := new(rest.Server)

	go func() {
		err = srv.Run(handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Printf("Fibonacci rest server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	time.Sleep(2 * time.Second)

	logrus.Print("App Shutting Down")

	// GraceFul shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found")
	}
}
