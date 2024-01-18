package models

import (
	"time"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/ch/chschema"
)

type Platform int32

const (
	ANDROID Platform = iota
	ANDROID_DEBUG
	IOS
	IOS_DEBUG
)

type OpenTarget int32

const (
	APP OpenTarget = iota
	SAVE_KEY
	EMULATE
	EDIT
	SHARE
	EXPERIMENTAL_FM
	EXPERIMENTAL_SCREENSTREAMING
	SHARE_SHORTLINK
	SHARE_LONGLINK
	SHARE_FILE
	SAVE_DUMP
	MFKEY32
	OPEN_NFC_DUMP_EDITOR
	OPEN_FAPHUB
	OPEN_FAPHUB_CATEGORY
	OPEN_FAPHUB_SEARCH
	OPEN_FAPHUB_APP
	INSTALL_FAPHUB_APP
	HIDE_FAPHUB_APP
)

type Open struct {
	ch.CHModel `ch:"table:open"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	Target OpenTarget
	Arg    *string
}

type FlipperGattInfo struct {
	ch.CHModel `ch:"table:flipper_gatt_info"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	FlipperVersion string
}

type FlipperRpcInfo struct {
	ch.CHModel `ch:"table:flipper_rpc_info"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	SdCardIsAvailable  bool
	InternalFreeBytes  int64
	InternalTotalBytes int64
	ExternalFreeBytes  int64
	ExternalTotalBytes int64
	FirmwareForkName   *string
	FirmwareGitUrl     *string
}

type SynchronizationEnd struct {
	ch.CHModel `ch:"table:synchronization_end"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	SubGHZCount           int32
	RFIDCount             int32
	NFCCount              int32
	InfraredCount         int32
	IButtonCount          int32
	SynchronizationTimeMs int64
	ChangesCount          int32
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

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	UpdateFrom   string
	UpdateTo     string
	UpdateId     int64
	UpdateStatus UpdateStatus
}

type UpdateFlipperStart struct {
	ch.CHModel `ch:"table:update_flipper_start"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	UpdateFrom string
	UpdateTo   string
	UpdateId   int64
}

type SubGhzProvisioning struct {
	ch.CHModel `ch:"table:subghz_provisioning"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	RegionNetwork  string
	RegionSim1     string
	RegionSim2     string
	RegionIp       string
	RegionSystem   string
	RegionProvided string
	IsRoaming      bool
}

type DebugInfo struct {
	ch.CHModel `ch:"table:debug_info_metric"`

	UUID        chschema.UUID
	Platform    Platform
	Time        time.Time `ch:",pk"`
	UserSession *chschema.UUID
	AppVersion  *string

	Key   string
	Value string
}
