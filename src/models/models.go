package models

import (
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/ch/chschema"
)

type Platform int32

const (
	ANDROID Platform = iota
	IOS
)

type AppOpen struct {
	ch.CHModel `ch:"table:app_open"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type ExperimentalOpenFileManager struct {
	ch.CHModel `ch:"table:experimental_open_fm"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type ExperimentalOpenScreenStreaming struct {
	ch.CHModel `ch:"table:experimental_open_screenstreaming"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type FlipperGattInfo struct {
	ch.CHModel `ch:"table:flipper_gatt_info"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`

	FlipperVersion string
}

type FlipperRpcInfo struct {
	ch.CHModel `ch:"table:flipper_rpc_info"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`

	SdCardIsAvailable  bool
	InternalFreeBytes  int64
	InternalTotalBytes int64
	ExternalFreeBytes  int64
	ExternalTotalBytes int64
}

type OpenEdit struct {
	ch.CHModel `ch:"table:open_edit"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type OpenEmulate struct {
	ch.CHModel `ch:"table:open_emulate"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type OpenSaveKey struct {
	ch.CHModel `ch:"table:open_save_key"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type OpenShare struct {
	ch.CHModel `ch:"table:open_share"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`
}

type SynchronizationEnd struct {
	ch.CHModel `ch:"table:synchronization_end"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`

	SubGHZCount           int32
	RFIDCount             int32
	NFCCount              int32
	InfraredCount         int32
	IButtonCount          int32
	SynchronizationTimeMs int64
}

type UpdateStatus int32

const (
	COMPLETED UpdateStatus = iota
	CANCELED
	FAILED_DOWNLOAD
	FAILED_PREPARE
	FAILED_UPLOAD
	FAILED
)

type UpdateFlipperEnd struct {
	ch.CHModel `ch:"table:update_flipper_end"`

	UUID     chschema.UUID
	Platform Platform
	Time     time.Time `ch:",pk"`

	UpdateFrom   string
	UpdateTo     string
	UpdateId     int64
	UpdateStatus UpdateStatus
}

type UpdateFlipperStart struct {
	ch.CHModel `ch:"table:update_flipper_start"`

	UUID       chschema.UUID
	Platform   Platform
	Time       time.Time `ch:",pk"`
	UpdateFrom string
	UpdateTo   string
	UpdateId   int64
}
