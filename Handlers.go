package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetHeaders(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Pokemon(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	fmt.Fprintln(w, "{\"Pokemon\": 1}")
}
func Trainer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Trainer")
}

func PokemonSpecific(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	pokeID := vars["pkmnID"]
	fmt.Fprintln(w, "{\"pmkn\": " + pokeID + "}")
}
