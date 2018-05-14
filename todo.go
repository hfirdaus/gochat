package main

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"name"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"due"`
	User	  string	`json:"user"`
}

type TodoDisplay struct {
	ID        int
	Task      string
	Completed bool
	DueDate   string
	User 	  string
}

type Page struct {
	Title	string
	Todos	[]TodoDisplay
	CUser 	string
}