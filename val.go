package goval

// Warper
type Val interface {
	Interface() interface{}
	Val() Val
	Type() Type

	// Value() (driver.Value, error) // implement database.sql.driver.Valuer
	// Scan(value interface{}) error // implement database.sql.Scanner
}
