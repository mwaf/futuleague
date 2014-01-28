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
	ident := mux.Vars(r)["player"]
	player, err := Player{}.FetchByIdentifier(ident)
	if err == nil {
		returnJson(w, player)
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Could not find player.")
	}
}
func (v1 APIv1) createPlayer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var player Player
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
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Player exists but could not be fetched.")
			return
		}
		v1.returnPlayerWithRedirect(w, r, player, http.StatusSeeOther)
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to create player.")
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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not marshal JSON.")
	}
}
