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
	rd := handlers.NewRedirection(l)
	sh := handlers.NewShortner(l)
	// initialized new servMux
	sm := http.NewServeMux()

	sm.Handle("/", rd)
	sm.Handle("/shorten",sh)
	http.ListenAndServe(":8080", sm)
}
