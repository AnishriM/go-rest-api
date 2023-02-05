package main

import (
	"net/http"

	"github.com/AnishriM/go-rest-api/internal/comment"
	"github.com/AnishriM/go-rest-api/internal/database"
	transportHTTP "github.com/AnishriM/go-rest-api/internal/transport/http"
	logrus "github.com/sirupsen/logrus"
)

// App - the struct which contains things like pointers to the database.
type App struct {
	Name    string
	Version string
}

// Run - Sets up our application
func (app App) Run() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(
		logrus.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		},
	).Info("Setting up an application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	if err := database.MigrateDB(db); err != nil {
		return err
	}
	commentService := comment.NewService(db)
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		logrus.Error("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	logrus.Info("GO REST APIs")
	app := App{
		Name:    "comment-service-application",
		Version: "1.1",
	}
	if err := app.Run(); err != nil {
		logrus.Error("Error occured while starting the application")
		logrus.Fatal(err)
	}
}
