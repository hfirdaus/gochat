package main

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type TodoDisplay struct {
	ID        int
	Name      string
	Completed bool
	Due       string
}

type Page struct {
	Title	string
	Todos	[]TodoDisplay
	Name 	string
}