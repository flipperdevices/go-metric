package main

import (
	"context"
	"github.com/flipperdevices/go-metric/internal/proto"
	"github.com/flipperdevices/go-metric/migrations"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/ch/chschema"
	"github.com/uptrace/go-clickhouse/chmigrate"
	"log"
	"time"
)

var db = ch.Connect(ch.WithDSN("clickhouse://localhost:19000/metric?sslmode=disable"))

func saveEvent(
	ctx context.Context,
	uid chschema.UUID,
	platform Platform,
	event *proto.MetricEventsCollection,
) error {
	var err error
	if event.GetAppOpen() != nil {
		_, err = db.NewInsert().Model(&AppOpen{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}).Exec(ctx)
	}

	if err != nil {
		return err
	}
	return nil
}

func applyMigration(db *ch.DB, ctx context.Context) {
	migrator := chmigrate.NewMigrator(db, migrations.Migrations)
	err := migrator.Init(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("End create migration")
}
