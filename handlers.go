package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/gorilla/mux"
	"html/template"
)

// https://golang.org/doc/articles/wiki/
// for later: https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Print(err)
		fmt.Fprintf(w,"An error has occurred.")
		return
	}
	t.Execute(w, Page{Title: "Home"})

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todos.html")
	if err != nil {
		fmt.Print(err)
		fmt.Fprintf(w,"An error has occurred.")
		return
	}
	t.Execute(w, Page{Title: "Todo List"})

	//if Username == "" {
	//	r.ParseForm()
	//	Username = r.FormValue("username")
	//}
	//fmt.Fprintf(w, "<h1>Hi %s!</h1>", Username)
	//
	//fmt.Fprintf(w, "<h1>Add a todo</h1>"+
	//	"<form action=\"/todos\" method=\"POST\">"+
	//	"<textarea name=\"name\"></textarea><br>"+
	//	"<input type=\"submit\" value=\"Save\">"+
	//	"</form>")
	//todos := FindAllTodos()
	//
	//fmt.Fprintf(w,"<ul>")
	//for i := 0; i < len(todos); i++ {
	//	fmt.Fprintf(w,"<li>%s</li>", todos[i].Name)
	//}
	//fmt.Fprintf(w,"</ul>")
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
	TodoIndex(w, r)
}