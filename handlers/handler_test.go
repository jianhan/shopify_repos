package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"path/filepath"

	"strings"

	"github.com/google/go-github/github"
	"github.com/jianhan/shopify_repos/api"
	"github.com/jianhan/shopify_repos/store"
	"github.com/sirupsen/logrus"
)

func TestShopify_index(t *testing.T) {
	path, _ := filepath.Abs("../")
	if err := os.Chdir(path); err != nil {
		panic(err)
	}
	testID1, testName1, testID2, testName2 := int64(1), "Test 1", int64(2), "Test 2"
	forkTrue, forkFalse := true, false
	twoDaysAgo := github.Timestamp{time.Now().Add(-time.Hour * 48)}
	oneDayAgo := github.Timestamp{time.Now().Add(-time.Hour * 24)}
	url1, url2 := "http://www.test1.com", "http://www.test2.com"
	expectedRepos := []*github.Repository{
		{
			ID:        &testID1,
			Name:      &testName1,
			Fork:      &forkTrue,
			UpdatedAt: &twoDaysAgo,
			Homepage:  &url1,
		},
		{
			ID:        &testID2,
			Name:      &testName2,
			Fork:      &forkFalse,
			UpdatedAt: &oneDayAgo,
			Homepage:  &url2,
		},
	}
	mockedShopifyAPI := &api.SPFGithubRepoFetcherMock{
		FetchFunc: func() ([]*github.Repository, error) {
			return expectedRepos, nil
		},
	}
	repoStoreMock := &store.RepoMock{
		GetReposFunc: func() []*github.Repository {
			return expectedRepos
		},
		SetReposFunc: func(items []*github.Repository) {

		},
		IsExpiredFunc: func() bool {
			return false
		},
	}
	shopify := &Shopify{log: logrus.New(), shopifyAPI: mockedShopifyAPI, repoStore: repoStoreMock}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shopify.index)
	handler.ServeHTTP(rr, req)
	body, _ := ioutil.ReadAll(rr.Result().Body)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// check html string contains the fields we expected
	if !strings.Contains(string(body), testName1) {
		t.Fatal("expected testName1 but not found")
	}
	if !strings.Contains(string(body), testName2) {
		t.Fatal("expected testName2 but not found")
	}
	if !strings.Contains(string(body), url1) {
		t.Fatal("expected url1 but not found")
	}
	if !strings.Contains(string(body), url2) {
		t.Fatal("expected url2 but not found")
	}
	if !strings.Contains(string(body), "Yes") {
		t.Fatal("expected Forked to be Yes but not found")
	}
	if !strings.Contains(string(body), "No") {
		t.Fatal("expected Forked to be No but not found")
	}
	if !strings.Contains(string(body), "2 days ago") {
		t.Fatal("expected 2 days ago but not found")
	}
	if !strings.Contains(string(body), "1 day ago") {
		t.Fatal("expected 1 day ago but not found")
	}
}
