package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipperdevices/go-metric/internal/proto"
	"github.com/flipperdevices/go-metric/internal/proto/events"
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
	var e any
	if event.GetAppOpen() != nil {
		e = AppOpen{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetExperimentalOpenFm() != nil {
		e = ExperimentalOpenFileManager{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetExperimentalOpenScreenstreaming() != nil {
		e = ExperimentalOpenScreenStreaming{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetFlipperGattInfo() != nil {
		e = FlipperGattInfo{
			UUID:           uid,
			Platform:       platform,
			Time:           time.Now(),
			FlipperVersion: event.GetFlipperGattInfo().FlipperVersion,
		}
	} else if event.GetFlipperRpcInfo() != nil {
		e = FlipperRpcInfo{
			UUID:               uid,
			Platform:           platform,
			Time:               time.Now(),
			SdCardIsAvailable:  event.GetFlipperRpcInfo().SdcardIsAvailable,
			InternalFreeBytes:  event.GetFlipperRpcInfo().InternalFreeByte,
			InternalTotalBytes: event.GetFlipperRpcInfo().InternalTotalByte,
			ExternalFreeBytes:  event.GetFlipperRpcInfo().ExternalFreeByte,
			ExternalTotalBytes: event.GetFlipperRpcInfo().ExternalTotalByte,
		}
	} else if event.GetOpenEdit() != nil {
		e = OpenEdit{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetOpenEmulate() != nil {
		e = OpenEmulate{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetOpenSaveKey() != nil {
		e = OpenSaveKey{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetOpenShare() != nil {
		e = OpenShare{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	} else if event.GetSynchronizationEnd() != nil {
		e = SynchronizationEnd{
			UUID:                  uid,
			Platform:              platform,
			Time:                  time.Now(),
			SubGHZCount:           event.GetSynchronizationEnd().SubghzCount,
			RFIDCount:             event.GetSynchronizationEnd().RfidCount,
			NFCCount:              event.GetSynchronizationEnd().NfcCount,
			InfraredCount:         event.GetSynchronizationEnd().InfraredCount,
			IButtonCount:          event.GetSynchronizationEnd().IbuttonCount,
			SynchronizationTimeMs: event.GetSynchronizationEnd().SynchronizationTimeMs,
		}
	} else if event.GetUpdateFlipperEnd() != nil {
		var status UpdateStatus
		switch event.GetUpdateFlipperEnd().UpdateStatus {
		case events.UpdateFlipperEnd_COMPLETED:
			status = COMPLETED
		case events.UpdateFlipperEnd_CANCELED:
			status = CANCELED
		case events.UpdateFlipperEnd_FAILED_DOWNLOAD:
			status = FAILED_DOWNLOAD
		case events.UpdateFlipperEnd_FAILED_PREPARE:
			status = FAILED_PREPARE
		case events.UpdateFlipperEnd_FAILED_UPLOAD:
			status = FAILED_UPLOAD
		case events.UpdateFlipperEnd_FAILED:
			status = FAILED
		}
		e = UpdateFlipperEnd{
			UUID:         uid,
			Platform:     platform,
			Time:         time.Now(),
			UpdateFrom:   event.GetUpdateFlipperEnd().UpdateFrom,
			UpdateTo:     event.GetUpdateFlipperEnd().UpdateTo,
			UpdateId:     event.GetUpdateFlipperEnd().UpdateId,
			UpdateStatus: status,
		}
	} else if event.GetUpdateFlipperStart() != nil {
		e = UpdateFlipperStart{
			UUID:       uid,
			Platform:   platform,
			Time:       time.Now(),
			UpdateFrom: event.GetUpdateFlipperStart().UpdateFrom,
			UpdateTo:   event.GetUpdateFlipperStart().UpdateTo,
			UpdateId:   event.GetUpdateFlipperStart().UpdateId,
		}
	} else {
		return errors.New("can't find table for event")
	}

	_, err := db.NewInsert().Model(&e).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

func applyMigration(ctx context.Context) {
	migrator := chmigrate.NewMigrator(db, migrations.Migrations)
	err := migrator.Init(ctx)

	if err != nil {
		panic(err)
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}
	if group.IsZero() {
		fmt.Println("there are no new migrations to run (database is up to date)")
	}

	log.Println("End create migration")
}
