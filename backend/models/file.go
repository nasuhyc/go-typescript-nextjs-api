package models

import (
	"time"
)

// ANCHOR - File model
type File struct {
	ID        uint64
	TableType string    `json:"-"`
	TableID   uint      `json:"-"`
	File      string    `json:"file"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
