package entities

import "time"

type Institution struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Abbreviation string     `json:"abbreviation"`
	City         string     `json:"city"`
	Address      string     `json:"address"`
	CreatedAt    *time.Time `json:"created_at"`
}
