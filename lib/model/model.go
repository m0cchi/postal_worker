package model

import "time"

type To struct {
	ModuleName string `json:"module_name"`
	Address    string `json:"address"`
}

// PostalMatter is ...
type PostalMatter struct {
	Body   string    `json:"body"`
	Format string    `json:"format"`
	To     []To      `json:"to"`
	From   string    `json:"from"`
	Date   time.Time `json:"date"`
}
