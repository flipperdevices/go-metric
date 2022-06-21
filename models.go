package main

import (
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/ch/chschema"
	"time"
)

type Platform int32

const (
	android Platform = 0
	ios              = 1
)

type AppOpen struct {
	ch.CHModel `ch:"table:app_open"`

	UUID     chschema.UUID `ch:",pk"`
	Platform Platform
	Time     time.Time
}
