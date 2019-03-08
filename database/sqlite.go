package database

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../models"
)

var dB *gorm.DB

func initDB() *gorm.DB {
	var err error
	dB, err = gorm.Open("sqlite3", "pr_api.db")

	if err != nil {
		panic("SHIT")
	}

	dB.AutoMigrate(&models.User{})

	return dB
}

func saveData(db *gorm.DB, data models.Query) {

	for _, stargazer := range data.Repository.Stargazers.Nodes {
		company := strings.ToLower(strings.Join(strings.Fields(stargazer.Company), ""))
		if company == "makeschool" || company == "@makeschool" {
			fmt.Println("found makeschool user:", stargazer.Login)
			user := &models.User{Login: stargazer.Login, PullRequests: make([]models.PullRequest, len(stargazer.PullRequests.Nodes))}
			for _, node := range stargazer.PullRequests.Nodes {
				fmt.Println("Found pr:", node.Title)
				pullRequest := &models.PullRequest{Title: node.Title, Date: node.CreatedAt}
				user.PullRequests = append(user.PullRequests, *pullRequest)
				fmt.Println("added:", user.PullRequests)
			}
			db.NewRecord(user)
			db.Create(&user)
		}
	}
}
