package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err == nil {

		for _, values := range books {
			if bookId == values.ID {

				json.NewEncoder(w).Encode(values)
			}

		}
	}
}

func GetPaths(w http.ResponseWriter, r *http.Request) {
	var infoAPI InfoAPI
	for _, values := range routes {

		infoAPI = append(infoAPI,
			Info{Path: values.Pattern, Description: values.Description})

	}

	json.NewEncoder(w).Encode(infoAPI)
}
