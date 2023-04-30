package models

type Url struct {
	ID     uint   `gorm:"primary_key"`
	Url    string `gorm:"column:url"`
	Hash   string `gorm:"column:hash"`
	UserID uint   `gorm:"column:user_id"`
}

func (u *Url) HasHash() bool {
	if u.Hash != "" {
		return true
	}

	return false
}

//

type UrlRequest struct {
	Url    string `json:"url"`
	Hash   string `json:"hash"`
	UserID uint   `json:"-"`
}

//

type UrlResponse struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}
