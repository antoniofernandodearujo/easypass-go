package models

// Ticket represents a ticket in the database.
type Ticket struct {
	ID           string `db:"id"`
	Name 		 string `db:"name"`
	Description  string `db:"description"`
	Date 	     string `db:"date"`
	Hour		 string `db:"hour"`
	Price        float64 `db:"price"`
	Quantity     int `db:"quantity"`
}
