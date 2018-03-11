package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jianhan/shopify_repos/api"
	"github.com/sirupsen/logrus"
)

// Shopify is a struct which have multiply handlers for different incomming requests.
type Shopify struct {
	log *logrus.Logger
	// depend on the interface, not implementation for testing.
	shopifyAPI api.SPFGithubRepoFetcher
}

// index is the main handler func which will fetch and display repos on HTML page.
func (s *Shopify) index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/shopify/index.html")
	if err != nil {
		s.log.Fatal(err)
		panic(err)
	}
	err = t.Execute(w, "TEST")
}

// Serve is the entry point of start up for HTTP server, which will be called by main func.
func Serve(addrAndPort string, shopifyAPI api.SPFGithubRepoFetcher) {
	shopify := &Shopify{log: logrus.New(), shopifyAPI: shopifyAPI}
	r := mux.NewRouter()
	r.HandleFunc("/", shopify.index)
	http.ListenAndServe(addrAndPort, r)
}
