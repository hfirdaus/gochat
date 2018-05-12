package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/gorilla/mux"
	"html/template"
	"time"
	"strconv"
)

// https://golang.org/doc/articles/wiki/
// for later: https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, Page{Title: "Home"})

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todos.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, Page{Title: "Todo List", Todos: FindAllTodos()})
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
	time, err := time.Parse(time.RFC3339, r.FormValue("due-date") + "T00:00:00Z")
	if err != nil {
		panic(err)
	}
	InsertTodo(Todo{Name: r.FormValue("name"), Due: time, Completed: false})
	TodoIndex(w, r)
}

func TodoComplete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	fmt.Print(id)
	if err != nil {
		panic(err)
	}
	ToggleTodoCompletedValue(id)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	fmt.Print(id)
	if err != nil {
		panic(err)
	}
	DeleteTodoById(id)
}