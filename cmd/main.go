package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kjg116/webserver-pyramid/internal/routers"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")

	router := routers.InitRoutes()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
