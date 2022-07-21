package models

import "time"

type Url struct {
	hash string
	url  string
	created_at time.Time
}
