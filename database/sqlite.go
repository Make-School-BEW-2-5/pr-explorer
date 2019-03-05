package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/Make-School-BEW-2-5/pr-explorer/models"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "pr_api.db")

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.PullRequest{}, &models.User{})

	return db, nil
}

// func SaveData(data models.Query) {
// 	db, err := Init()

// 	for stargazer := range data.Repository.Stargazers {
// 		company := strings.ToLower(strings.Join(strings.Fields(Stargazer.Company), ""))
// 		if company == "makeschool" {
// 			user := &models.User{Login: stargazer.Login, PullRequests: make([]models.PullRequest, len(stargazer.Nodes.PullRequests.Nodes))}
// 			for node := range stargazer.Nodes.PullRequests.Nodes {
// 				pullRequest := &models.PullRequest{Title: node.Title, Date: node.createdAt}
// 				db.NewRecord(pullRequest)
// 				db.Create(&pullRequest)
// 				append(user.PullRequests, pullRequest)
// 			}
// 			db.NewRecord(user)
// 			db.Create(&user)
// 		}
// 	}
// }
