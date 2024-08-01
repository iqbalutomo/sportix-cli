package main

import (
	"log"
	"sportix-cli/config"
	"sportix-cli/internal/cli"
	"sportix-cli/internal/db"
	handler "sportix-cli/internal/handler/user"
	repository "sportix-cli/internal/repository/user"

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

	app := tview.NewApplication()
	cli.MainCLI(app, cli.Handler{
		User: userHandler,
	})

}
