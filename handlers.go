package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/gorilla/mux"
)

// https://golang.org/doc/articles/wiki/

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Add a todo</h1>"+
		"<form action=\"/\" method=\"POST\">"+
		"<textarea name=\"name\"></textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(FindAllTodos()); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", todoId, todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := InsertTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func TodoSave(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("NAME => %s\n", r.FormValue("name"))
	InsertTodo(Todo{Name: r.FormValue("name")})
}