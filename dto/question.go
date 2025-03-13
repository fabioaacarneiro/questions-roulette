package dto

import (
	"time"
)

type Question struct {
	ID        *int       `json:"id,omitempty"`
	Question  string     `json:"Question"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
