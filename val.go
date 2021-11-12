package goval

import "database/sql/driver"

// Warper
type Val interface {
	Interface() interface{}
	Type() Type

	// Value implementation driver.Valuer interface.
	Value() (driver.Value, error)

	// Scan implementation sql.Scanner interface.
	Scan(src interface{}) error
}

type Type interface {
}

const (
	ValTypeNil = iota
	ValTypeBool
	ValTypeInt8
	ValTypeInt16
	ValTypeInt32
	ValTypeInt64
	ValTypeUint8
	ValTypeUint16
	ValTypeUint32
	ValTypeUint64
	ValTypeFloat32
	ValTypeFloat64
	ValTypeString
	ValTypeBytes
	ValTypeTime
	ValTypeDate32
	ValTypeDate64
)
