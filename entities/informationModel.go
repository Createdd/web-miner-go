package entities

import (
	"time"
)

type author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// InformationModel with defined structure
type InformationModel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    author    `json:"author"`
	Published time.Time `json:"updated_at"`
}
