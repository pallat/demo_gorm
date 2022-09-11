package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("fail to connect database test.db: %s\n", err)
	}

	db.Migrator().DropTable(&Todo{})
	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Println(err)
	}

	if r := db.Create(&Todo{Title: "Learn Go", Done: true}); r.Error != nil {
		log.Println(r.Error)
	}
	if r := db.Create(&Todo{Title: "Practice TDD", Done: false}); r.Error != nil {
		log.Println(r.Error)
	}
	if r := db.Create(&Todo{Title: "Make an API", Done: false}); r.Error != nil {
		log.Println(r.Error)
	}
	if r := db.Create(&Todo{Title: "Shop a book", Done: false}); r.Error != nil {
		log.Println(r.Error)
	}

	var todos []Todo
	r := db.Find(&todos)
	if err := r.Error; err != nil {
		log.Panic(err)
	}

	for _, todo := range todos {
		fmt.Printf("%d %s, status(%t)\n", todo.ID, todo.Title, todo.Done)
	}
}

type Todo struct {
	gorm.Model
	Title string
	Done  bool
}
