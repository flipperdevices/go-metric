package main

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/flipperdevices/go-metric/src/models"
	"github.com/uptrace/go-clickhouse/ch"
	"log"
	"net/http"

	"github.com/flipperdevices/go-metric/src/reporter"
	"github.com/flipperdevices/go-metric/src/repository"
)

func main() {
	ctx := context.Background()

	var cfg models.Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}
	db := ch.Connect(
		ch.WithAddr(cfg.DBAddr),
		ch.WithDatabase(cfg.DBName),
		ch.WithInsecure(true),
	)
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	repo := repository.New(db)

	// db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))

	if err := repo.ApplyMigration(ctx); err != nil {
		panic(err)
	}

	report := reporter.New(repo)

	http.HandleFunc("/report", report.Report)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
