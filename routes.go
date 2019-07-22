package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Description string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes

func init() {
	routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			"Index",
			Index,
		},
		Route{
			"TodoShow",
			"GET",
			"/books/{bookId}",
			"Get specific book with ID",
			GetBook,
		},
		Route{
			"TodoIndex",
			"GET",
			"/books",
			"Get all Books",
			GetBooks,
		},

		Route{
			"info",
			"GET",
			"/info",
			"Get Paths info",
			GetPaths,
		},
	}
}
