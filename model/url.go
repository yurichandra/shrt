package model

import "time"

// URL represent URL model.
type URL struct {
	ID          uint
	OriginalURL string
	ShortURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
