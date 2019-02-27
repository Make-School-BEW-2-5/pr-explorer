package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// GetData sends a GraphQL request to GitHub and returns a collection of users and their pull requests
func GetData(c echo.Context) error {
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
					PullRequests struct {
						Nodes []struct {
							CreatedAt time.Time
							Author    struct {
								Login string
							}
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
	return c.JSON(http.StatusOK, query.Repository.Stargazers.Nodes)
	// fmt.Println("	CreatedAt:", query.Viewer.CreatedAt)

}
