package domain

import "github.com/google/uuid"

type Shortener struct {
	ID   uint      `gorm:"primaryKey"`
	Url  string    `gorm:"column:url"`
	Hash string    `gorm:"column:hash"`
	User uuid.UUID `gorm:"foreignKey:UUID"`
}

type ShortenerResponse struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func (s *Shortener) toResponse() ShortenerResponse {
	return ShortenerResponse{
		Url: s.Url,
		Hash: s.Hash,
	}
}