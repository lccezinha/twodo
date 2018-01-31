package main

import (
	"fmt"

	"github.com/lccezinha/twodo/cmd/environment"
)

func main() {
	app := environment.Init()
	todos, err := app.ListService.Run()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("List all", todos)
	fmt.Println("Create single Todo")
	err = app.CreateService.Run("Title", "Some Description")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Done....")

	todos, err = app.ListService.Run()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("List all")
	for _, todo := range todos {
		fmt.Println(todo.ID, todo.Title, todo.Description, todo.CreatedAt, todo.Done)
	}
}
