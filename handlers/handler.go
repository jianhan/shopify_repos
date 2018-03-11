package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Shopify struct {
	log *logrus.Logger
}

func (s *Shopify) index(w http.ResponseWriter, r *http.Request) {
}

func Serve(addrAndPort string) {
	shopify := &Shopify{log: logrus.New()}
	r := mux.NewRouter()
	r.HandleFunc("/", shopify.index)
	http.ListenAndServe(addrAndPort, r)
}
