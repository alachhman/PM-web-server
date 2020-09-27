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

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func Pokemon(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	fmt.Fprintln(w, "{\"Pokemon\": "+strconv.Itoa(findCount("pokemon"))+"}")
}
func Trainer(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	trainerName := vars["trainerName"]
	fmt.Fprintln(w, findData(trainerName, "trainers"))
}

func PokemonSpecific(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	pkmnName := vars["pkmnName"]
	fmt.Fprintln(w, findData(pkmnName, "pokemon"))
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
