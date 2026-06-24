package entities

import "time"

type Institution struct {
	ID           int64
	Name         string
	Abbreviation string
	City         string
	Address      string
	CreatedAt    time.Time
}
