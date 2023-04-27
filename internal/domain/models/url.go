package models

type Url struct {
	ID     uint   `json:"id"`
	Url    string `json:"url"`
	Hash   string `json:"hash"`
	UserID uint   `json:"-"`
}
