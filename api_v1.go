package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type APIv1 struct{}

var v1 = APIv1{} // no need to have several of these

func routeAPIv1(r *mux.Router) {
	// TODO subrouter for headers matcing json + apiv1
	r.HandleFunc("/root", v1.root)
	r.HandleFunc("/{game}", v1.game)
}

func (v1 APIv1) root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Games games games, i.e. FIFA, NHL.")
}

func (v1 APIv1) game(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game := vars["game"]
	fmt.Fprintf(w, "info for game, get them clubs for %s", game)
}

type Game struct {
	Name  string `json:"name"`
	Clubs []Club `json:"clubs"`
}

type Club struct {
	Name   string  `json:"name"`
	League string  `json:"name"`
	Stars  float64 `json:"stars"`
}
