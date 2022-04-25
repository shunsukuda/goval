package goval

// Warper
type Val interface {
	Type() Type

	// Value() (driver.Value, error) // implement database.sql.driver.Valuer
	// Scan(value interface{}) error // implement database.sql.Scanner
}
