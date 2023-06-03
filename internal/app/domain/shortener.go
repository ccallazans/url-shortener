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

func (s *Shortener) ToResponse() ShortenerResponse {
	return ShortenerResponse{
		Url:  s.Url,
		Hash: s.Hash,
	}
}

func ShortenersToResponse(shorteners []Shortener) []ShortenerResponse {

	var shortenersReponse []ShortenerResponse
	for _, shortener := range shorteners {
		shortenersReponse = append(shortenersReponse, shortener.ToResponse())
	}

	return shortenersReponse
}
