package database

import (
	"encoding/json"
	"log"
	"strings"

	"../models"
	bolt "go.etcd.io/bbolt"
)

func Init() *bolt.DB {
	db, err := bolt.Open("pr_api.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SaveData(db *bolt.DB, query models.Query) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}

		targetCompany := "makeschool"

		for _, stargazer := range query.Repository.Stargazers.Nodes {
			company := strings.ToLower(strings.Join(strings.Fields(stargazer.Company), ""))
			if company == targetCompany || company == "@"+targetCompany {
				user := &models.User{
					Login:        stargazer.Login,
					PullRequests: make([]models.PullRequest, len(stargazer.PullRequests.Nodes)),
				}
				for _, pr := range stargazer.PullRequests.Nodes {
					pullRequest := &models.PullRequest{
						Title: pr.Title,
						Date:  pr.CreatedAt,
					}
					user.PullRequests = append(user.PullRequests, *pullRequest)
				}
				encodedUser, err := json.Marshal(user)
				if err != nil {
					return err
				}
				b.Put([]byte(user.Login), encodedUser)
			}
		}
		return nil
	})
}
