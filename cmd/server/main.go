package main

import (
	"log"

	"github.com/amirt713/finance-app/cmd"
	"github.com/amirt713/finance-app/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	app := cmd.Application()
	PORT := config.LoadEnv("PORT")

	if err := app.Listen(PORT); err != nil {
		log.Fatal(err)

	}

}
