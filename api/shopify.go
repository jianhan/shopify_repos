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
	Org    string
	Client *github.Client
	Opts   *github.RepositoryListByOrgOptions
}

// Fetch gets all repos from shopify via github developer API.
func (s *Shopify) Fetch() ([]*github.Repository, error) {
	// Do not need pagination, etc.. therefore, omit response
	repos, _, err := s.Client.Repositories.ListByOrg(context.Background(), s.Org, s.Opts)
	if err != nil {
		return nil, err
	}
	return repos, err
}
