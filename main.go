package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
	"os"

	"github.com/Ujk768/url-shortener/handlers"
)

func main() {
	l := log.New(os.Stdout, "url-shortener", log.LstdFlags)
	//initialzed new handler
	sh := handlers.NewRedirection(l)
	// initialized new servMux
	sm := http.NewServeMux()

	sm.Handle("/", sh)

	http.ListenAndServe(":8080", sm)
}
