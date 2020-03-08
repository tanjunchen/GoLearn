package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const tokenv4 = "be22f4726892ae65f21f3d8f3832540c157e993f"

func main() {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tokenv4},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}
	fmt.Println("    Login:", query.Viewer.Login)
	fmt.Println("CreatedAt:", query.Viewer.CreatedAt)

}
