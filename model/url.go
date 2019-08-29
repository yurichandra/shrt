package model

import "time"

// URL represent URL model.
type URL struct {
	ID          uint
	UserID      uint
	OriginalURL string
	ShortURL    string
	ExpiredDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	User        User
}
