package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/sobocinski/api-go-update/db"
	"github.com/sobocinski/api-go-update/infrastructure/pgsql"
	"github.com/sobocinski/api-go-update/internal/car"
	transportHTTP "github.com/sobocinski/api-go-update/internal/transport/http"
	"github.com/sobocinski/api-go-update/internal/user"
	"net/http"
	"os"
)



func init() {
	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
}

type App struct {
	Name    string
	Version string
}



func (app *App) Run() error {
	log.Println("Starting app...")

	//db config
	ctx := context.Background()
	db, err := db.NewDatabase(ctx)
	if err != nil {
		log.Panic("Database error", err)
	}
	defer db.Close()

	userRepository := pgsql.NewUserRepository(db)
	userService := user.NewService(userRepository)

	carRepository := pgsql.NewCarRepository(db)
	carService := car.NewService(carRepository)

	handler := transportHTTP.NewHandler(userService, carService)
	handler.SetupRoutes()

	log.Info("Listen at :8080")
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}



func main() {
	app := App{"API", "0.1"}
	app.Run()
}
