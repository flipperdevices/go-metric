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

func (r *Repository) CheckDb(
	ctx context.Context,
) (bool, error) {
	var num int
	if err := r.db.QueryRowContext(ctx, "SELECT 1").Scan(&num); err != nil {
		return false, err
	}
	return num == 1, nil
}

func (r *Repository) SaveEvent(
	ctx context.Context,
	uid chschema.UUID,
	sid *chschema.UUID,
	version *string,
	platform models.Platform,
	event *proto.MetricEventsCollection,
) error {
	q := r.db.NewInsert()
	switch unpackedEvent := event.GetEvent().(type) {
	case *proto.MetricEventsCollection_Open:
		var target models.OpenTarget
		switch unpackedEvent.Open.Target {
		case events.Open_APP:
			target = models.APP
		case events.Open_SAVE_KEY:
			target = models.SAVE_KEY
		case events.Open_EMULATE:
			target = models.EMULATE
		case events.Open_EDIT:
			target = models.EDIT
		case events.Open_SHARE:
			target = models.SHARE
		case events.Open_EXPERIMENTAL_FM:
			target = models.EXPERIMENTAL_FM
		case events.Open_EXPERIMENTAL_SCREENSTREAMING:
			target = models.EXPERIMENTAL_SCREENSTREAMING
		case events.Open_SHARE_SHORTLINK:
			target = models.SHARE_SHORTLINK
		case events.Open_SHARE_LONGLINK:
			target = models.SHARE_LONGLINK
		case events.Open_SHARE_FILE:
			target = models.SHARE_FILE
		case events.Open_SAVE_DUMP:
			target = models.SAVE_DUMP
		case events.Open_MFKEY32:
			target = models.MFKEY32
		case events.Open_OPEN_NFC_DUMP_EDITOR:
			target = models.OPEN_NFC_DUMP_EDITOR
		}

		q = q.Model(&models.Open{
			UUID:        uid,
			Platform:    platform,
			Time:        time.Now(),
			UserSession: sid,
			AppVersion:  version,
			Target:      target,
		})
	case *proto.MetricEventsCollection_FlipperGattInfo:
		q = q.Model(&models.FlipperGattInfo{
			UUID:           uid,
			Platform:       platform,
			Time:           time.Now(),
			UserSession:    sid,
			AppVersion:     version,
			FlipperVersion: unpackedEvent.FlipperGattInfo.FlipperVersion,
		})
	case *proto.MetricEventsCollection_FlipperRpcInfo:
		var forkName *string = nil
		if len(unpackedEvent.FlipperRpcInfo.FirmwareForkName) > 0 {
			forkName = &unpackedEvent.FlipperRpcInfo.FirmwareForkName
		}
		var gitUrl *string = nil
		if len(unpackedEvent.FlipperRpcInfo.FirmwareGitUrl) > 0 {
			gitUrl = &unpackedEvent.FlipperRpcInfo.FirmwareGitUrl
		}
		q = q.Model(&models.FlipperRpcInfo{
			UUID:               uid,
			Platform:           platform,
			Time:               time.Now(),
			UserSession:        sid,
			AppVersion:         version,
			SdCardIsAvailable:  unpackedEvent.FlipperRpcInfo.SdcardIsAvailable,
			InternalFreeBytes:  unpackedEvent.FlipperRpcInfo.InternalFreeByte,
			InternalTotalBytes: unpackedEvent.FlipperRpcInfo.InternalTotalByte,
			ExternalFreeBytes:  unpackedEvent.FlipperRpcInfo.ExternalFreeByte,
			ExternalTotalBytes: unpackedEvent.FlipperRpcInfo.ExternalTotalByte,
			FirmwareForkName:   forkName,
			FirmwareGitUrl:     gitUrl,
		})
	case *proto.MetricEventsCollection_SynchronizationEnd:
		q = q.Model(&models.SynchronizationEnd{
			UUID:                  uid,
			Platform:              platform,
			Time:                  time.Now(),
			UserSession:           sid,
			AppVersion:            version,
			SubGHZCount:           unpackedEvent.SynchronizationEnd.SubghzCount,
			RFIDCount:             unpackedEvent.SynchronizationEnd.RfidCount,
			NFCCount:              unpackedEvent.SynchronizationEnd.NfcCount,
			InfraredCount:         unpackedEvent.SynchronizationEnd.InfraredCount,
			IButtonCount:          unpackedEvent.SynchronizationEnd.IbuttonCount,
			SynchronizationTimeMs: unpackedEvent.SynchronizationEnd.SynchronizationTimeMs,
			ChangesCount:          unpackedEvent.SynchronizationEnd.ChangesCount,
		})
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
		q = q.Model(&models.UpdateFlipperEnd{
			UUID:         uid,
			Platform:     platform,
			Time:         time.Now(),
			UserSession:  sid,
			AppVersion:   version,
			UpdateFrom:   unpackedEvent.UpdateFlipperEnd.UpdateFrom,
			UpdateTo:     unpackedEvent.UpdateFlipperEnd.UpdateTo,
			UpdateId:     unpackedEvent.UpdateFlipperEnd.UpdateId,
			UpdateStatus: status,
		})
	case *proto.MetricEventsCollection_UpdateFlipperStart:
		q = q.Model(&models.UpdateFlipperStart{
			UUID:        uid,
			Platform:    platform,
			Time:        time.Now(),
			UserSession: sid,
			AppVersion:  version,
			UpdateFrom:  unpackedEvent.UpdateFlipperStart.UpdateFrom,
			UpdateTo:    unpackedEvent.UpdateFlipperStart.UpdateTo,
			UpdateId:    unpackedEvent.UpdateFlipperStart.UpdateId,
		})
	case *proto.MetricEventsCollection_SubghzProvisioning:
		q = q.Model(&models.SubGhzProvisioning{
			UUID:           uid,
			Platform:       platform,
			Time:           time.Now(),
			UserSession:    sid,
			AppVersion:     version,
			RegionNetwork:  unpackedEvent.SubghzProvisioning.RegionNetwork,
			RegionSim1:     unpackedEvent.SubghzProvisioning.RegionSim_1,
			RegionSim2:     unpackedEvent.SubghzProvisioning.RegionSim_2,
			RegionIp:       unpackedEvent.SubghzProvisioning.RegionIp,
			RegionSystem:   unpackedEvent.SubghzProvisioning.RegionSystem,
			RegionProvided: unpackedEvent.SubghzProvisioning.RegionProvided,
			IsRoaming:      unpackedEvent.SubghzProvisioning.IsRoaming,
		})

	default:
		return errors.New("can't find table for event")
	}

	if _, err := q.Exec(ctx); err != nil {
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
