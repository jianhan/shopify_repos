package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jianhan/shopify_repos/api"
	"github.com/jianhan/shopify_repos/store"
	"github.com/sirupsen/logrus"
)

// Shopify is a struct which have multiply handlers for different incomming requests.
type Shopify struct {
	log *logrus.Logger
	// depend on the interface, not implementation for testing.
	shopifyAPI api.SPFGithubRepoFetcher
	repoStore  store.Repo
}

// index is the main handler func which will fetch and display repos on HTML page.
func (s *Shopify) index(w http.ResponseWriter, r *http.Request) {
	// check if cache expired, if it is then fetch API and load into cache
	if s.repoStore.IsExpired() {
		repos, err := s.shopifyAPI.Fetch()
		if err != nil {
			s.log.Fatal(err)
			panic(err)
		}
		s.repoStore.SetRepos(repos)
		s.log.WithField("repos_count", len(repos)).Info("Finish fetching repos from github API")
	}
	repos := s.repoStore.GetRepos()
	s.log.WithField("repos_count", len(repos)).Info("Finish fetching repos from in memory cache")
	t, err := template.New("index.html").Funcs(template.FuncMap{
		"readableTime":    UTCToLocal,
		"readableBoolean": ReadableBoolean,
	}).ParseFiles("views/shopify/index.html")
	if err != nil {
		s.log.Fatal(err)
		panic(err)
	}
	err = t.Execute(w, repos)
	if err != nil {
		s.log.Fatal(err)
		panic(err)
	}
}

// Serve is the entry point of start up for HTTP server, which will be called by main func.
func Serve(addrAndPort string, shopifyAPI api.SPFGithubRepoFetcher, repoStore store.Repo) {
	shopify := &Shopify{log: logrus.New(), shopifyAPI: shopifyAPI, repoStore: repoStore}
	r := mux.NewRouter()
	r.HandleFunc("/", shopify.index)
	http.ListenAndServe(addrAndPort, r)
}
