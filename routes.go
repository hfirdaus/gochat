package main

import (
	"net/http"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoSave",
		"POST",
		"/save",
		TodoSave,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoComplete",
		"POST",
		"/complete",
		TodoComplete,
	},
	Route{
		"TodoDelete",
		"POST",
		"/delete",
		TodoDelete,
	},
	Route{
		"Name",
		"POST",
		"/name",
		Name,
	},
}