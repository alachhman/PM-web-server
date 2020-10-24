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
		"Home",
		"GET",
		"/",
		Home,
	},
	Route{
		"Favicon",
		"GET",
		"/favicon.ico",
		Favicon,
	},
	Route{
		"FrontEndCss",
		"GET",
		"/static/css/{fileName}",
		handleFrontEndCss,
	},
	Route{
		"FrontEndJs",
		"GET",
		"/static/js/{fileName}",
		handleFrontEndJs,
	},
	Route{
		"FrontEndMedia",
		"GET",
		"/static/media/{fileName}",
		handleFrontEndMedia,
	},
	Route{
		"TrainerList",
		"GET",
		"/trainer/list",
		TrainerList,
	},
	Route{
		"FrontEndFile",
		"GET",
		"/{fileName}",
		handleFrontEndFile,
	},
	Route{
		"PokemonCount",
		"GET",
		"/pkmn",
		PokemonCount,
	},
	Route{
		"Trainer",
		"GET",
		"/trainer/{trainerName}",
		Trainer,
	},
	Route{
		"TrainerImage",
		"GET",
		"/trainer/image/{imageFileName}",
		TrainerImage,
	},
}
