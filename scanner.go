package goval

import (
	"errors"
	"time"
)

func scanVal(src interface{}) (Val, error) {
	switch cv := src.(type) {
	case int64:
		val := (*Int64)(&cv)
		return val, nil
	case float64:
		val := (*Float64)(&cv)
		return val, nil
	case bool:
		return Bool(cv), nil
	case []byte:
		return Bytes(cv), nil
	case string:
		return String(cv), nil
	case time.Time:
		return Time(cv), nil
	case nil:
		return nil, nil
	default:
		return nil, errors.New("sql scan error")
	}
}
