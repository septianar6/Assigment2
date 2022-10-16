package main

import (
	"assigment2/controller"
	"assigment2/database"
	"assigment2/routes"
	"fmt"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("error start database", err)
		return
	}

	ctl := controller.New(db)

	err = routes.StartServer(ctl)
	if err != nil {
		fmt.Println("error start server", err)
		return
	}
}
