package entities

import "time"

type Event struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	InstitutionId int64     `json:"institution_id"`
	CreatedAt     time.Time `json:"created_at"`
}
