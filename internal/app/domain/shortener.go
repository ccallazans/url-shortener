package domain

type Shortener struct {
	ID   uint   `gorm:"primary_key"`
	Url  string `gorm:"column:url"`
	Hash string `gorm:"column:hash"`
	User User   `gorm:"embedded"`
}
