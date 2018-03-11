package main

import (
	"context"

	"os"

	"github.com/google/go-github/github"
	"github.com/jianhan/shopify_repos/api"
	"github.com/jianhan/shopify_repos/handlers"
	"github.com/jianhan/shopify_repos/store"
	"golang.org/x/oauth2"
)

func main() {
	spfFetcher, repoStore := bootstrap()
	handlers.Serve(os.Getenv("ADDRESS_PORT"), spfFetcher, repoStore)
}

// bootstrap initializes all dependencies we needed in main func.
func bootstrap() (api.SPFGithubRepoFetcher, store.Repo) {
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, nil)
	client := github.NewClient(tc)
	return api.NewShopify(os.Getenv("REPO_USER"), client, nil), store.NewRepoStore()
}
