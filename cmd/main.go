package main

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/services"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/handlers/http"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/pubsub"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/repositories"
	"log"
)

func run() error {
	config, err := internal.NewConfig()
	if err != nil {
		return err
	}

	kafka := pubsub.NewKafka()

	repo, err := repositories.NewPostgres(config.DSN)
	if err != nil {
		return err
	}

	companyService := services.NewCompanyService(*repo, kafka)

	handler := http.New(config.AppPort, config.AppMode, companyService)
	return handler.Run()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("error runnning the app: %s", err)
	}
}
