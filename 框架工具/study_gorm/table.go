package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model

	String  string
	Bool    bool
	Byte    byte
	Int     int
	Int64   int64
	Uint    uint
	Uint64  uint64
	Float64 float64
	Time    time.Time

	StringPtr  *string
	BoolPtr    *bool
	BytePtr    *byte
	IntPtr     *int
	Int64Ptr   *int64
	UintPtr    *uint
	Uint64Ptr  *uint64
	Float64Ptr *float64
	TimePtr    *time.Time

	StringNull  sql.NullString
	BoolNull    sql.NullBool
	ByteNull    sql.NullByte
	Int64Null   sql.NullInt64
	Float64Null sql.NullFloat64
	TimeNull    sql.NullTime
}
