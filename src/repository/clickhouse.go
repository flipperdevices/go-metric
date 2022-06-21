package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/ch/chschema"
	"github.com/uptrace/go-clickhouse/chmigrate"

	"github.com/flipperdevices/go-metric/internal/proto"
	"github.com/flipperdevices/go-metric/internal/proto/events"
	"github.com/flipperdevices/go-metric/migrations"
	"github.com/flipperdevices/go-metric/src/models"
)

type Repository struct {
	db *ch.DB
}

func New(db *ch.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveEvent(
	ctx context.Context,
	uid chschema.UUID,
	platform models.Platform,
	event *proto.MetricEventsCollection,
) error {
	var e any
	switch unpackedEvent := event.GetEvent().(type) {
	case *proto.MetricEventsCollection_AppOpen:
		e = models.AppOpen{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_ExperimentalOpenFm:
		e = models.ExperimentalOpenFileManager{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_ExperimentalOpenScreenstreaming:
		e = models.ExperimentalOpenScreenStreaming{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_FlipperGattInfo:
		e = models.FlipperGattInfo{
			UUID:           uid,
			Platform:       platform,
			Time:           time.Now(),
			FlipperVersion: unpackedEvent.FlipperGattInfo.FlipperVersion,
		}
	case *proto.MetricEventsCollection_FlipperRpcInfo:
		e = models.FlipperRpcInfo{
			UUID:               uid,
			Platform:           platform,
			Time:               time.Now(),
			SdCardIsAvailable:  unpackedEvent.FlipperRpcInfo.SdcardIsAvailable,
			InternalFreeBytes:  unpackedEvent.FlipperRpcInfo.InternalFreeByte,
			InternalTotalBytes: unpackedEvent.FlipperRpcInfo.InternalTotalByte,
			ExternalFreeBytes:  unpackedEvent.FlipperRpcInfo.ExternalFreeByte,
			ExternalTotalBytes: unpackedEvent.FlipperRpcInfo.ExternalTotalByte,
		}
	case *proto.MetricEventsCollection_OpenEdit:
		e = models.OpenEdit{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_OpenEmulate:
		e = models.OpenEmulate{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_OpenSaveKey:
		e = models.OpenSaveKey{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_OpenShare:
		e = models.OpenShare{
			UUID:     uid,
			Platform: platform,
			Time:     time.Now(),
		}
	case *proto.MetricEventsCollection_SynchronizationEnd:
		e = models.SynchronizationEnd{
			UUID:                  uid,
			Platform:              platform,
			Time:                  time.Now(),
			SubGHZCount:           unpackedEvent.SynchronizationEnd.SubghzCount,
			RFIDCount:             unpackedEvent.SynchronizationEnd.RfidCount,
			NFCCount:              unpackedEvent.SynchronizationEnd.NfcCount,
			InfraredCount:         unpackedEvent.SynchronizationEnd.InfraredCount,
			IButtonCount:          unpackedEvent.SynchronizationEnd.IbuttonCount,
			SynchronizationTimeMs: unpackedEvent.SynchronizationEnd.SynchronizationTimeMs,
		}
	case *proto.MetricEventsCollection_UpdateFlipperEnd:
		var status models.UpdateStatus
		switch unpackedEvent.UpdateFlipperEnd.UpdateStatus {
		case events.UpdateFlipperEnd_COMPLETED:
			status = models.COMPLETED
		case events.UpdateFlipperEnd_CANCELED:
			status = models.CANCELED
		case events.UpdateFlipperEnd_FAILED_DOWNLOAD:
			status = models.FAILED_DOWNLOAD
		case events.UpdateFlipperEnd_FAILED_PREPARE:
			status = models.FAILED_PREPARE
		case events.UpdateFlipperEnd_FAILED_UPLOAD:
			status = models.FAILED_UPLOAD
		case events.UpdateFlipperEnd_FAILED:
			status = models.FAILED
		}
		e = models.UpdateFlipperEnd{
			UUID:         uid,
			Platform:     platform,
			Time:         time.Now(),
			UpdateFrom:   unpackedEvent.UpdateFlipperEnd.UpdateFrom,
			UpdateTo:     unpackedEvent.UpdateFlipperEnd.UpdateTo,
			UpdateId:     unpackedEvent.UpdateFlipperEnd.UpdateId,
			UpdateStatus: status,
		}
	case *proto.MetricEventsCollection_UpdateFlipperStart:
		e = models.UpdateFlipperStart{
			UUID:       uid,
			Platform:   platform,
			Time:       time.Now(),
			UpdateFrom: unpackedEvent.UpdateFlipperStart.UpdateFrom,
			UpdateTo:   unpackedEvent.UpdateFlipperStart.UpdateTo,
			UpdateId:   unpackedEvent.UpdateFlipperStart.UpdateId,
		}
	default:
		return errors.New("can't find table for event")
	}

	if _, err := r.db.NewInsert().Model(&e).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *Repository) ApplyMigration(ctx context.Context) error {
	migrator := chmigrate.NewMigrator(r.db, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return fmt.Errorf("Repository.ApplyMigration: %w", err)
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("Repository.ApplyMigration: %w", err)
	}
	if group.IsZero() {
		fmt.Println("there are no new migrations to run (database is up to date)")
	}

	log.Println("End create migration")
	return nil
}

func (r *Repository) CreateTables(ctx context.Context) error {
	if _, err := r.db.NewCreateTable().Model((*models.AppOpen)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.ExperimentalOpenFileManager)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.ExperimentalOpenScreenStreaming)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.FlipperGattInfo)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.FlipperRpcInfo)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.OpenEdit)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.OpenEmulate)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.OpenSaveKey)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.OpenShare)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.SynchronizationEnd)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.UpdateFlipperEnd)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := r.db.NewCreateTable().Model((*models.UpdateFlipperStart)(nil)).Exec(ctx); err != nil {
		return err
	}
	return nil
}
