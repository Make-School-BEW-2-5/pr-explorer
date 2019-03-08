package models

import (
	"time"
)

type PullRequest struct {
	Title  string
	Date   time.Time
	UserID uint
}
