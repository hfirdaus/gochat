package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var currentId int = 0

func init() {
	db, err := gorm.Open("postgres", "host=172.17.0.2 port=5432 user=gorm dbname=todo password=gorm")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Todo{})

	// Initial todos
	Insert(db, Todo{Name: "Thing 1", Completed:false, Due: time.Now()})
	Insert(db, Todo{Name: "Thing 2", Completed:true, Due: time.Now()})
}

func ModifyDatabase(f Modify, t Todo) Todo {
	db, err := gorm.Open("postgres", "host=172.17.0.2 port=5432 user=gorm dbname=todo password=gorm")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	f(db, t)
	return t
}

func SearchDatabase(f Find, Id int) {
	db, err := gorm.Open("postgres", "host=172.17.0.2 port=5432 user=gorm dbname=todo password=gorm")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	f(db, Id)
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
