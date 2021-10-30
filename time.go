package goval

import "time"

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
