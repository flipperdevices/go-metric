package main

import (
	"context"
	"log"
	"net/http"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"

	"github.com/flipperdevices/go-metric/src/reporter"
	"github.com/flipperdevices/go-metric/src/repository"
)

func main() {
	ctx := context.Background()

	db := ch.Connect(ch.WithDSN("clickhouse://localhost:19000/metric?sslmode=disable"))
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	repo := repository.New(db)

	db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))
	if err := repo.CreateTables(ctx); err != nil {
		panic(err)
	}

	//if err := repo.ApplyMigration(ctx); err != nil {
	//	panic(err)
	//}

	report := reporter.New(repo)

	http.HandleFunc("/report", report.Report)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
