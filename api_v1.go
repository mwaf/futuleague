package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type APIv1 struct{}

var v1 = APIv1{} // no need to have several of these

const (
	ACCEPT_HEADER = "application/vnd.futuleague.v1+json"
)

func routeAPIv1(r *mux.Router) {

	// POST routes with proper header
	post := r.Methods("POST").Headers("Accept", ACCEPT_HEADER).Subrouter()
	post.HandleFunc("/players", v1.createPlayer)

	// GET routes with proper headers
	get := r.Methods("GET").Headers("Accept", ACCEPT_HEADER).Subrouter()

	get.HandleFunc("/players/{player}", v1.player)
	get.HandleFunc("/players", v1.players)
	get.HandleFunc("/clubs", v1.clubs)

	// GET routes without proper headers but .json in the path instead
	getDotJson := r.Methods("GET").Subrouter()
	getDotJson.HandleFunc("/players/{player}.json", v1.player)
	getDotJson.HandleFunc("/players.json", v1.players)

	getDotJson.HandleFunc("/clubs.json", v1.clubs)
}


func (v1 APIv1) clubs(w http.ResponseWriter, r *http.Request) {
	clubs, err := Club{}.FetchAll()
	if err == nil {
		returnJson(w, clubs)
	} else {
		returnErrorJson(w, http.StatusBadRequest, "Could not fetch clubs.", err)
	}
}

func (v1 APIv1) player(w http.ResponseWriter, r *http.Request) {
	ident := mux.Vars(r)["player"]
	player, err := Player{}.FetchByIdentifier(ident)
	if err == nil {
		returnJson(w, player)
	} else {
		returnErrorJson(w, http.StatusBadRequest, "Could not find player.", err)
	}
	// sorry, being lazy and not differentiating between a DB
	// failure (Internal Server Error) and player not found (
}

func (v1 APIv1) players(w http.ResponseWriter, r *http.Request) {
	players, _ := Player{}.FetchAll()
	// ignoring error, just showing empty user list if it failed
	// (or is actually empty)
	returnJson(w, players)
}

func (v1 APIv1) createPlayer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	player := Player{Played: 0, Wins: 0, Losses: 0, Ties: 0}
	player.Identifier = r.Form.Get("identifier")
	player.Name = r.Form.Get("name")
	player.Rating = DEFAULT_PLAYER_RATING

	err := player.Save()
	switch {
	case err == nil:
		v1.returnPlayerWithRedirect(w, r, player, http.StatusCreated)
	case err.Error() == PLAYER_ALREADY_EXISTS_ERROR:
		player, err = Player{}.FetchByIdentifier(player.Identifier)
		if err != nil {
			returnErrorJson(w, http.StatusInternalServerError, "Player exists but could not be fetched.", err)
			return
		}
		v1.returnPlayerWithRedirect(w, r, player, http.StatusSeeOther)
	case err != nil:
		returnErrorJson(w, http.StatusInternalServerError, "Unable to create player.", err)
	}
}

func (v1 APIv1) returnPlayerWithRedirect(w http.ResponseWriter, r *http.Request, player Player, statusCode int) {
	header := w.Header()
	header.Add("Content-Type", "application/json")
	path := fmt.Sprintf("/players/%s", player.Identifier)
	http.Redirect(w, r, path, statusCode)
	output, err := json.Marshal(player)
	if err == nil {
		w.Write(output)
	}
	// If marshaling fails at this point it's still better to
	// return the redirect to the actual resource without a body
	// than confuse the client with an internal srever error (the
	// player was successfully created after all)
}

func returnJson(w http.ResponseWriter, v interface{}) {
	result, err := json.Marshal(v)
	if err == nil {
		header := w.Header()
		header.Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		returnErrorJson(w, http.StatusInternalServerError, "Could not marshal JSON.", err)
	}
}

func returnErrorJson(w http.ResponseWriter, status int, msg string, err error) {
	v := RootError{Error: JsonError{msg, err.Error()}}
	result, _ := json.Marshal(v)
	// ignore error here, we're returning one after all :)

	header := w.Header()
	header.Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}
