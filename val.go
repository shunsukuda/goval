package goval

// Warper
type Val interface {
	Interface() interface{}
	Type() Type
}

type Type int

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
)
