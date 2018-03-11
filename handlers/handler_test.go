package handlers

import (
	"net/http"
	"testing"

	"github.com/jianhan/shopify_repos/api"
	"github.com/jianhan/shopify_repos/store"
)

func TestShopify_index(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		s    *Shopify
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.index(tt.args.w, tt.args.r)
		})
	}
}

func TestServe(t *testing.T) {
	type args struct {
		addrAndPort string
		shopifyAPI  api.SPFGithubRepoFetcher
		repoStore   store.Repo
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Serve(tt.args.addrAndPort, tt.args.shopifyAPI, tt.args.repoStore)
		})
	}
}
