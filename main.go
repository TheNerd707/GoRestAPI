package main

import (
	"fmt"
	"net/http"

	"github.com/TheNerd707/GoRestAPI/internal/server"
)

func main() {
	fmt.Println("Starting server process")
	app := server.New()

	app.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	app.Route("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test")
	})

	app.Route("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Testttttt")
	})

	err := app.Listen(":3000", MuxHandler)

	if err != nil {
		panic(err)
	}
}


func MuxHandler(mux *http.ServeMux) {

}

