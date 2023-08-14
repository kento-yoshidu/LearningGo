package main

import (
	"fmt"
	"rest-api/db"
	"rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Success Migration!")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
