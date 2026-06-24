package entities

import "time"

type User struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Address       string    `json:"address"`
	Document      string    `json:"document"`
	Cellphone     string    `json:"cellphone"`
	Role          string    `json:"role"`
	InstitutionID *int64    `json:"institution_id"`
	Score         int16     `json:"score"`
	CreatedAt     time.Time `json:"created_at"`
	BirthDate     time.Time `json:"birthDate"`
}
