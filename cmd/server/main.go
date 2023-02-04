package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/AnishriM/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers to the database.
type App struct{}

// Run - Sets up our application
func (app App) Run() error {
	fmt.Println("Running Application")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	println("GO REST APIs")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error occured while starting the application", err.Error())
	}
}
