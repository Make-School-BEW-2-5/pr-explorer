package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Make-School-BEW-2-5/pr-explorer/models"
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

	err = client.Query(context.Background(), &models.Query, nil)
	if err != nil {
		log.Fatal(err)
	}
	return query
	// fmt.Println("	CreatedAt:", query.Viewer.CreatedAt)

}
