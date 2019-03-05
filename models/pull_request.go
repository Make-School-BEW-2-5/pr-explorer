package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PullRequest struct {
	gorm.Model

	Title string
	Date  time.Time
}
