package main

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/jianhan/shopify_repos/api"
	"github.com/jianhan/shopify_repos/handlers"
	"golang.org/x/oauth2"
)

func main() {
	// Get fetcher
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		// TODO: following token needed to be in ENV
		&oauth2.Token{AccessToken: "48e65f4e0070a283ed925f2b62501e7a75fc8966"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// TODO: options is nil now, may needed to be changes later if pagination needed to be implemented.
	spfFetcher := api.NewShopify("Shopify", client, nil)

	// TODO: address and port should be in ENV
	handlers.Serve("127.0.0.1:8080", spfFetcher)
}
