package main

import (
	"fmt"
	"github.com/google/go-github/v29/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
)

const token = "be22f4726892ae65f21f3d8f3832540c157e993f"

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	options := &github.PullRequestListOptions{State: "closed"}

	// list all repositories for the authenticated user
	repos, _, err := client.PullRequests.List(ctx, "istio", "istio.io", options)
	if err != nil {
		log.Printf("fail", err)
	}
	for i, _ := range repos {
		fmt.Println(i+1)
		//if *pr.Merged {
		//	fmt.Println(*pr.URL, *pr.State, *pr.Assignee, pr.Labels, pr.MergedAt, pr.MergedBy, pr.Number, pr.User)
		//}
	}
}
