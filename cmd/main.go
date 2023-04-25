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
	cfg, err := internal.NewConfig()
	if err != nil {
		return err
	}

	kafka := pubsub.NewKafka()

	repo, err := repositories.New(cfg.DSN, cfg.DBMaxPoolSize, cfg.DBConnAttempts, cfg.DBConnTimeoutSeconds)
	if err != nil {
		return err
	}

	if cfg.DBApplyMigrations == 1 {
		if err := repo.Migrate(); err != nil {
			return err
		}
	}

	companyService := services.NewCompanyService(*repo, kafka)

	handler := http.New(cfg.AppPort, cfg.AppMode, companyService)
	return handler.Run()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("error runnning the app: %s", err)
	}
}
