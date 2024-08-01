package main

import (
	"log"
	"sportix-cli/config"
	"sportix-cli/internal/cli"
	"sportix-cli/internal/db"
	"sportix-cli/internal/handler"
	"sportix-cli/internal/repository"

	"github.com/rivo/tview"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userHandler := handler.NewUserHandler(userRepo)

	fieldRepo := repository.NewFieldRepo(db)
	fieldHandler := handler.NewFieldHandler(fieldRepo)

	app := tview.NewApplication()
	cli.MainCLI(app, cli.Handler{
		User:  userHandler,
		Field: fieldHandler,
	})

}
