package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Shopify struct {
	log *logrus.Logger
}

func (s *Shopify) index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/shopify/index.html")
	if err != nil {
		s.log.Fatal(err)
		panic(err)
	}
	err = t.Execute(w, "TEST")
}

func Serve(addrAndPort string) {
	shopify := &Shopify{log: logrus.New()}
	r := mux.NewRouter()
	r.HandleFunc("/", shopify.index)
	http.ListenAndServe(addrAndPort, r)
}
