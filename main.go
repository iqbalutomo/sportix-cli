package main

import (
	"fmt"
	"log"
	"sportix-cli/config"
	"sportix-cli/internal/db"
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

	fmt.Println("database connected successfully!")
}
