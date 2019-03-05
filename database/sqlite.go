package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/Make-School-BEW-2-5/pr-explorer/models/"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "pr_api.db")

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.PullRequest, &models.User{})

	return db, nil
}
