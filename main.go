package main

import (
	"net/http"
	"os"

	"practice/server/handlers"
	"practice/server/store"
	"practice/server/utils"
)

func main() {

	// initialize the logger
	utils.NewLogger("server: ")

	//create a new store
	err := store.NewStore()
	if err != nil {
		utils.Logger.Error("Error connecting to database ", "error", err.Error())
		os.Exit(1)
	}

	http.HandleFunc("GET /task", handlers.GetTasks)
	http.HandleFunc("POST /createtask", handlers.CreateTask)
	http.HandleFunc("GET /task/{id}", handlers.GetTask)
	http.HandleFunc("DELETE /deletetask/{id}", handlers.DeleteTask)
	http.HandleFunc("GET /tag/{tagname}", handlers.GetTag)
	http.HandleFunc("GET /due/{yy}/{mm}/{dd}", handlers.GetTasksDue)

	err = http.ListenAndServe(":8080", nil)
	utils.Logger.Info("shutting down server ", "error", err.Error())
}
