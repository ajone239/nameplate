package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajone239/nameplate/internal/app"
)

const staticDir = "./web/build"

func main() {
	apiApp := app.NewApp(12, staticDir)

	routes := apiApp.Routes()

	fmt.Println("Listening on: localhost:80")

	err := http.ListenAndServe(":80", routes)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Closing up shop")
	}
}
