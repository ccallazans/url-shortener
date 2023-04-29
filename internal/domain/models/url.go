package models

type Url struct {
	ID     uint
	Url    string
	Hash   string
	UserID uint
}

//

type UrlRequest struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

//

type UrlResponse struct {
	Url      string `json:"url"`
	Hash     string `json:"hash"`
	Username string `json:"username"`
}
