package models

import "time"

type Session struct {
	UserName     string
	LastActivity time.Time
}
