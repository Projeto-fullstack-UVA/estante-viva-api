package entities

import "time"

type Event struct {
	ID           int64
	Name         string
	Description  string
	Location     string
	InstitutionId int64
	Created_at   time.Time
}
