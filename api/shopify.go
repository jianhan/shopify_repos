package api

import (
	"context"

	"github.com/google/go-github/github"
)

// SPFGithubRepoFetcher is a single method interface for fetching repos of shopify.
type SPFGithubRepoFetcher interface {
	Fetch() ([]*github.Repository, error)
}

// Shopify implements SPFGithubRepoFetcher.
type Shopify struct {
	org    string
	client *github.Client
	opts   *github.RepositoryListByOrgOptions
}

// Fetch gets all repos from shopify via github developer API.
func (s *Shopify) Fetch() ([]*github.Repository, error) {
	// Do not need pagination, etc.. therefore, omit response, might need later on.
	repos, _, err := s.client.Repositories.ListByOrg(context.Background(), s.org, s.opts)
	if err != nil {
		return nil, err
	}
	return repos, err
}

// NewShopify receives parameters for creating a api.Shopify and return an interface.
func NewShopify(org string, client *github.Client, opts *github.RepositoryListByOrgOptions) SPFGithubRepoFetcher {
	return &Shopify{
		org:    org,
		client: client,
		opts:   opts,
	}
}
