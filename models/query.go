package models

import "time"

type Query struct {
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
