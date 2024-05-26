package main

import (
	"context"
	"exchanger_test/internal/config"
	"exchanger_test/internal/handlers"
	"exchanger_test/internal/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logLevel := config.GetLoggerLevel()
	host := config.GetHostConfig()
	port := config.GetPortConfig()

	exchangerService := service.NewExchangerService()
	exchangeHandler := handlers.NewExchangerHandler(exchangerService)

	logrus.SetLevel(logLevel)
	server := http.Server{
		Addr: host + port,
	}

	http.Handle("/exchanger", exchangeHandler.GetExchanger())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	select {
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
		}

		close(done)
	}
}
