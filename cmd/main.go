package main

import (
	"log"

	"github.com/dimaglushkov/epam-xm-test-assignment/internal"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/services"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/events"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/handlers"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/repositories"
)

func run() error {
	cfg, err := internal.NewConfig()
	if err != nil {
		return err
	}

	repo, err := repositories.NewPostgres(cfg.DSN, cfg.DBMaxPoolSize, cfg.DBConnAttempts, cfg.DBConnTimeoutSeconds)
	if err != nil {
		return err
	}
	defer repo.Pool.Close()

	if cfg.DBApplyMigrations == 1 {
		if err := repo.Migrate(); err != nil {
			return err
		}
	}

	kafka, err := events.NewKafkaWriter(cfg.KafkaBrokers, cfg.KafkaTopic)
	if err != nil {
		return err
	}
	defer kafka.Close()

	companyService := services.NewCompanyService(cfg.AppName, repo, kafka)

	handler := handlers.NewHTTPHandler(cfg.AppPort, cfg.AppMode, cfg.AppSignKey, companyService)

	return handler.Run()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("error runnning the app: %s", err)
	}
}
