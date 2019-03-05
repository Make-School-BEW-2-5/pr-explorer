package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// GetData sends a GraphQL request to GitHub and returns a collection of users and their pull requests
func GetData() interface{} {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Repository struct {
			Stargazers struct {
				Nodes []struct {
					Login        string
					Company      string
					PullRequests struct {
						Nodes []struct {
							CreatedAt time.Time
							Title     string
						}
					} `graphql:"pullRequests(first: 100)"`
				}
			} `graphql:"stargazers(first: 100)"`
		} `graphql:"repository(name: \"iruka\", owner: \"iruka-dev\")"`
	}

	err = client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Fatal(err)
	}
	return query
	// fmt.Println("	CreatedAt:", query.Viewer.CreatedAt)

}
