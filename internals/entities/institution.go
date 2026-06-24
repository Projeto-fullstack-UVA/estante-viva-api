package entities

import "time"

type Intitutions struct {
	ID           int64
	Name         string
	Abbreviation string
	City         string
	Address      string
	created_at   time.Time
}
