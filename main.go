package main

import (
	"arh/api"
	"arh/cmd"
	"arh/pkg/config"
	"os"
)

func main() {

	if config.MODE == "DEV" {
		cmd.Execute()
	} else if config.MODE == "PROD" {
		port := os.Getenv("PORT")

		if port == "" {
			port = "21500"
		}
		app := api.AppSchema{}
		app.Initialize()
		app.Run(":" + port)
	}
}
