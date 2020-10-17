package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Pokemon",
		"GET",
		"/pkmn",
		Pokemon,
	},
	Route{
		"Trainer",
		"GET",
		"/trainer/{trainerName}",
		Trainer,
	},
	Route{
		"TrainerList",
		"GET",
		"/trainer",
		TrainerList,
	},
	Route{
		"pkmn",
		"GET",
		"/pkmn/{pkmnName}",
		PokemonSpecific,
	},
}
