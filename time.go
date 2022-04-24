package goval

/*
import (
	"time"

	"github.com/apache/arrow/go/arrow"
)

const (
	Hour = time.Hour
	Day  = 24 * Hour
)

type Time time.Time

func (e Time) Time() time.Time {
	return time.Time(e)
}

func (e Time) Interface() interface{} {
	return time.Time(e)
}

func (e Time) Type() Type {
	return ValTypeTime
}

func (e Time) ToBool() Bool {
	cv, _ := e.ToBoolCheck()
	return cv
}

func (e Time) ToBoolCheck() (Bool, Err) {
	return Bool(!e.Time().IsZero()), nil
}

func (e Time) ToTime() Time {
	return e
}

func (e Time) ToTimeCheck() (Time, Err) {
	return e, nil
}

func (e Time) ToArrowDate32() arrow.Date32 {
	return arrow.Date32(e.Time().Unix() / int64(Day))
}

func (e Time) ToArrowDate64() arrow.Date64 {
	return arrow.Date64(e.Time().Unix() / int64(Day))
}

func (e Time) ToArrowTime32() arrow.Time32 {
	return arrow.Time32(e.Time().Unix() % int64(Day))
}

func (e Time) ToArrowTime64() arrow.Time64 {
	return arrow.Time64(e.Time().Unix() % int64(Day))
}

func (e Time) ToArrowTimestamp(tu arrow.TimeUnit) arrow.Timestamp {
	switch tu {
	case arrow.Nanosecond:
		return arrow.Timestamp(e.Time().Local().UnixNano())
	case arrow.Microsecond:
		return arrow.Timestamp(e.Time().Local().UnixMicro())
	case arrow.Millisecond:
		return arrow.Timestamp(e.Time().Local().UnixMilli())
	case arrow.Second:
		return arrow.Timestamp(e.Time().Local().Unix())
	default:
		return arrow.Timestamp(e.Time().Local().UnixNano())
	}
}

func (e Time) ToArrowTypeTimestamp(tu arrow.TimeUnit) arrow.TimestampType {
	return arrow.TimestampType{
		Unit:     tu,
		TimeZone: e.Time().Location().String(),
	}
}

func (e Time) ArrowTypeDurationType(tu arrow.TimeUnit) arrow.DurationType {
	return arrow.DurationType{
		Unit: tu,
	}
}
*/
