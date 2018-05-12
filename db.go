package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"time"
)

// Reference: https://www.codementor.io/engineerapart/getting-started-with-postgresql-on-mac-osx-are8jcopb

func init() {
	db := openTodoDb()
	defer db.Close()

	if !db.HasTable(&Todo{}) {
		db.AutoMigrate(&Todo{})
		InsertTodo(Todo{Name: "Thing 1", Completed: false, Due: time.Now()})
		InsertTodo(Todo{Name: "Thing 2", Completed: true, Due: time.Now()})
	}
}

func openTodoDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func InsertTodo(t Todo) Todo {
	db := openTodoDb()
	defer db.Close()
	t.ID = time.Now().Nanosecond()
	db.Create(&t)
	return t
}

func DeleteTodoById(Id int) {
	var todo Todo
	db := openTodoDb()
	defer db.Close()
	db.First(&todo, Id)
	db.Delete(todo)
}

func ToggleTodoCompletedValue(Id int) {
	var todo Todo
	db := openTodoDb()
	defer db.Close()
	db.First(&todo, Id)
	db.Model(todo).Update("Completed", !todo.Completed)
}


func UpdateTodo(t Todo) {
	db := openTodoDb()
	defer db.Close()
	db.Model(t).Update("Name", t.Name)
	db.Model(t).Update("Completed", t.Completed)
	db.Model(t).Update("Due", t.Due)
}

func FindTodoById(Id int) Todo {
	var todo Todo
	db := openTodoDb()
	defer db.Close()
	db.First(todo, Id)
	return todo
}

func FindAllTodos() []Todo {
	var todos []Todo
	db := openTodoDb()
	defer db.Close()
	db.Find(&todos)
	return todos
}
