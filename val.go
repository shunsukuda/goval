package goval

// Warper
type Val interface {
	Interface() interface{}
	Val() Val
	Type() Type

	// Value implementation driver.Valuer interface.
	//Value() (driver.Value, error)

	// Scan implementation sql.Scanner interface.
	//Scan(src interface{}) error
}
