package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"time"
)

var currentId int = 0

func init() {
	db := openTodoDb()
	defer db.Close()

	fmt.Println(db.HasTable(&Todo{}))

	db.AutoMigrate(&Todo{})

	Insert(db, Todo{Name: "Thing 1", Completed: false, Due: time.Now()})
	Insert(db, Todo{Name: "Thing 2", Completed: true, Due: time.Now()})
}

func openTodoDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func ModifyDatabase(f Modify, t Todo) Todo {
	db := openTodoDb()
	defer db.Close()
	f(db, t)
	return t
}

type Modify func(db *gorm.DB, t Todo)

func Insert(db *gorm.DB, t Todo) {
	currentId++
	t.Id = currentId
	db.Create(t)
}

func Delete(db *gorm.DB, t Todo) {
	db.Delete(t)
}

func Update(db *gorm.DB, t Todo) {
	db.Model(t).Update("Name", t.Name)
	db.Model(t).Update("Completed", t.Completed)
	db.Model(t).Update("Due", t.Due)
}

type Find func(db *gorm.DB, Id int)

func FindById(db *gorm.DB, Id int) Todo {
	var todo Todo
	db.First(todo, Id)
	return todo
}

func FindAllTodos() []Todo {
	var todos []Todo
	db := openTodoDb()
	db.Find(&todos)
	return todos
}
