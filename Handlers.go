package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Pokemon(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	fmt.Fprintln(w, "{\"Pokemon\": 1}")
}
func Trainer(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	trainerName := vars["trainerName"]
	fmt.Fprintln(w, fileWalk(trainerName, "trainers"))
}

func PokemonSpecific(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	pkmnName := vars["pkmnName"]
	fmt.Fprintln(w, fileWalk(pkmnName, "pokemon"))
}

func fileWalk(pkmnName string, datatype string) string {
	var files []string
	var out string = "["
	var isFirst bool = true
	root := "data/" + datatype
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
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
	}
	out += "]"
	return out
}
