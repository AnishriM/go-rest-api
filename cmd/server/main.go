package main

import "fmt"

// App - the struct which contains things like pointers to the database.
type App struct{}

// Run - Sets up our application
func (app App) Run() error {
	fmt.Println("Running Application")
	return nil
}

func main() {
	println("GO REST APIs")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error occured while starting the application", err.Error())
	}
}
