package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
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

func Name(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Username = r.FormValue("username")
	http.Redirect(w, r, "/todos", 301)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todos.html")
	if err != nil {
		panic(err)
	}

	todos := FindAllTodos()
	todoDisplays := make([]TodoDisplay, len(todos))
	for i := 0; i < len(todos); i++ {
		if todos[i].Due.IsZero() {
			todoDisplays[i] = TodoDisplay{	Name:todos[i].Name,
				ID: todos[i].ID,
				Completed: todos[i].Completed}
		} else {
			dueDate := todos[i].Due.Format("Mon Jan 2")
			today := time.Now().Truncate(24*time.Hour)
			someDay := todos[i].Due.Truncate(24*time.Hour)

			if today.Equal(someDay) {
				dueDate = "Today"
			} else if today.Add(24*time.Hour).Equal(someDay) {
				dueDate = "Tomorrow"
			} else if someDay.Add(24*time.Hour).Equal(today) {
				dueDate = "Yesterday"
			}
			if someDay.Before(today) && !todos[i].Completed {
				dueDate += " (Missed!)"
			}

			todoDisplays[i] = TodoDisplay{	Name:todos[i].Name,
				ID: todos[i].ID,
				Completed: todos[i].Completed,
				Due: dueDate}
		}
	}

	t.Execute(w, Page{Title: "Todo List", Todos: todoDisplays, Name: Username})
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
	name := r.FormValue("name")
	dueDateRaw := r.FormValue("due-date")
	if dueDateRaw != "" {
		dueDate, err := time.Parse(time.RFC3339, dueDateRaw + "T00:00:00Z")
		if err != nil {
			panic(err)
		}
		InsertTodo(Todo{Name: name, Due: dueDate, Completed: false})
	} else {
		InsertTodo(Todo{Name: name, Completed: false})
	}
	http.Redirect(w, r, "/todos", 301)
}

func TodoComplete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	fmt.Print(id)
	if err != nil {
		panic(err)
	}
	ToggleTodoCompletedValue(id)
	http.Redirect(w, r, "/todos", 301)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	fmt.Print(id)
	if err != nil {
		panic(err)
	}
	DeleteTodoById(id)
	http.Redirect(w, r, "/todos", 301)
}