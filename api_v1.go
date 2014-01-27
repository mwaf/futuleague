package main

import (
	"encoding/json"
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
	games := Game{}.FetchAll()
	gameList := make([]string, len(games))
	for i, g := range games {
		gameList[i] = g.Name
	}
	returnJson(w, gameList)
}

func (v1 APIv1) game(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["game"]
	game, err := Game{}.FetchByName(name)
	if err == nil {
		returnJson(w, game)
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Could not find game.")
	}
}

func returnJson(w http.ResponseWriter, v interface{}) {
	result, err := json.Marshal(v)
	if err == nil {
		header := w.Header()
		header.Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Could not marshal JSON.")
	}
}
