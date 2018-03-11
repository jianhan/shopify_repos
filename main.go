package main

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/jianhan/shopify_repos/api"
	"github.com/jianhan/shopify_repos/handlers"
	"golang.org/x/oauth2"
)

func main() {
	spfFetcher := bootstrap()
	handlers.Serve("127.0.0.1:8080", spfFetcher)
}

// bootstrap initializes all dependencies we needed in main func.
func bootstrap() api.SPFGithubRepoFetcher {
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, nil)
	client := github.NewClient(tc)
	// TODO: options is nil now, may needed to be changes later if pagination needed to be implemented.
	return api.NewShopify("Shopify", client, nil)
}
