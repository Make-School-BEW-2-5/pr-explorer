package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Login        string
	Refer        uint
	PullRequests []PullRequest
}
