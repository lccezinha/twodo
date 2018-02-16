package main

import (
	"log"
	"net/http"

	"github.com/lccezinha/twodo/cmd/web"
)

func main() {
	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/todos", web.CreateHandler)
	http.HandleFunc("/destroy", web.DestroyHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Print("Running server... http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Some shit happen to server: %s", err.Error())
	}

	// app := environment.Init()
	// todos, err := app.ListService.Run()
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// fmt.Println("List all", todos)
	// fmt.Println("Create single Todo")
	// err = app.CreateService.Run("Title", "Some Description")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("Done....")
	//
	// todos, err = app.ListService.Run()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("List all")
	// for _, todo := range todos {
	// 	fmt.Println(todo.ID, todo.Title, todo.Description, todo.CreatedAt, todo.Done)
	// }
}
