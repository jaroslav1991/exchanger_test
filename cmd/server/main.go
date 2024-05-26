package main

import (
	"context"
	"exchanger_test/internal/handlers"
	"exchanger_test/internal/service"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	exchangerService := service.NewExchangerService()
	exchangeHandler := handlers.NewExchangerHandler(exchangerService)

	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/exchanger", exchangeHandler.GetExchanger())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	select {
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		close(done)
	}
}
