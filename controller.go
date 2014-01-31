package main

import (
	"encoding/json"
	"net/http"
)

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

func returnWithRedirect(w http.ResponseWriter, r *http.Request, v interface{}, path string, statusCode int) {
	header := w.Header()
	header.Add("Content-Type", "application/json")
	http.Redirect(w, r, path, statusCode)
	output, err := json.Marshal(v)
	if err == nil {
		w.Write(output)
	}
	// If marshaling fails at this point it's still better to
	// return the redirect to the actual resource without a body
	// than confuse the client with an internal srever error (the
	// player was successfully created after all)
}
