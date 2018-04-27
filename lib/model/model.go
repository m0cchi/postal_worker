package model

import "time"

// PostalMatter is ...
type PostalMatter struct {
	ModuleName string    `json:"module_name"`
	Body       string    `json:"body"`
	Format     string    `json:"format"`
	To         string    `json:"to"`
	From       string    `json:"from"`
	Date       time.Time `json:"date"`
}
