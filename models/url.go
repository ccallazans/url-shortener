package models

import "time"

type UrlShort struct {
	ID          int       `json:"id"`
	Hash        string    `json:"hash"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}
