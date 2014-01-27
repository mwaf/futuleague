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
	get := r.Methods("GET").Subrouter()
	post := r.Methods("POST").Subrouter()

	get.HandleFunc("/players/{player}", v1.player)
	post.HandleFunc("/players", v1.createPlayer)

	get.HandleFunc("/root", v1.root)
	get.HandleFunc("/{game}", v1.game)
}

func (v1 APIv1) root(w http.ResponseWriter, r *http.Request) {
	games := Game{}.FetchAll()
	gameList := make([]string, len(games))
	for i, g := range games {
		gameList[i] = g.Name
	}
	returnJson(w, Root{Games: gameList})
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

func (v1 APIv1) player(w http.ResponseWriter, r *http.Request) {
}
func (v1 APIv1) createPlayer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var player Player
	player.Identifier = r.Form.Get("identifier")
	player.Name = r.Form.Get("name")
	player.Rating = DEFAULT_PLAYER_RATING
	fmt.Println("got error you! ", player.Save())

	path := fmt.Sprintf("/players/%s", player.Identifier)
	http.Redirect(w, r, path, http.StatusCreated)
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
