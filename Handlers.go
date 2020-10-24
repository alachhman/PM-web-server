package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func setHeaders(w http.ResponseWriter, contentType string){
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

func PokemonCount(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, "application/json; charset=UTF-8")
	fmt.Fprintln(w, "{\"Pokemon\": "+strconv.Itoa(findCount("pokemon"))+"}")
}

func Trainer(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	trainerName := vars["trainerName"]
	fmt.Fprintln(w, findData(trainerName, "trainers"))
}

func TrainerList(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, "application/json; charset=UTF-8")
	fmt.Fprintln(w, getTrainerList())
}

func TrainerImage(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, "image/png")
	vars := mux.Vars(r)
	fileName := vars["imageFileName"]
	http.ServeFile(w, r , "data/trainerImages/" + fileName)
}

func Home(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, "text/html; charset=UTF-8")
	http.ServeFile(w, r, "frontEnd/homepage/build/index.html")
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, "image/x-icon")
	http.ServeFile(w, r, "favicon.ico")
}

func handleFrontEndCss(w http.ResponseWriter, r *http.Request){
	setHeaders(w, "text/css")
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	http.ServeFile(w, r, "frontEnd/homepage/build/static/css/" + fileName)
}

func handleFrontEndJs(w http.ResponseWriter, r *http.Request){
	setHeaders(w,"application/javascript")
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	http.ServeFile(w, r, "frontEnd/homepage/build/static/js/" + fileName)
}

func handleFrontEndMedia(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	http.ServeFile(w, r, "frontEnd/homepage/build/static/media/" + fileName)
}

func handleFrontEndFile(w http.ResponseWriter, r *http.Request){
	setHeaders(w,"application/javascript")
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	http.ServeFile(w, r, "frontEnd/homepage/build/" + fileName)
}

func findCount(dataType string) int {
	var files []string
	var count int
	root := "data/" + dataType
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, _ = range files {
		count++
	}
	return count
}

func getTrainerList() string {
	var files []string
	var out = "{\"trainers\" : ["
	var isFirst = true
	root := "data/trainers"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if isFirst {
			isFirst = false;
		} else {
			content, err := ioutil.ReadFile(file)

			if err != nil {
				log.Fatal(err)
			}

			out += string(content) + ","
		}
	}
	out = strings.TrimRight(out, ",")
	out += "]}"
	return out
}

func findData(pkmnName string, datatype string) string {
	var files []string
	var out = ""
	if datatype == "pokemon" {
		out = "["
	}
	var isFirst = true
	root := "data/" + datatype
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if datatype == "pokemon" {
			if strings.Contains(file, pkmnName) {
				content, err := ioutil.ReadFile(file)

				if err != nil {
					log.Fatal(err)
				}

				if isFirst {
					out += string(content)
					isFirst = false
				} else {
					out += ", " + string(content)
				}
			}
		} else {
			if file == "data\\trainers\\" + pkmnName + ".json" {
				content, err := ioutil.ReadFile(file)

				if err != nil {
					log.Fatal(err)
				}

				if isFirst {
					out += string(content)
					isFirst = false
				} else {
					out += ", " + string(content)
				}
			}
		}
	}
	if datatype == "pokemon" {
		out += "]"
	} else {
		out = strings.Replace(out, "}", "", 1)
		out += ",\n\"pokemonData\": " + findData(pkmnName, "pokemon") + " }"
	}

	return out
}
