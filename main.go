package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello Go...")
	http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
		reqBody, err := io.ReadAll(rq.Body)
		if err != nil {
			fmt.Println("Error in handleFunc")
		}
		fmt.Fprintf(rw, "Hi from HandleFunc %s", reqBody)
	})
	// sm := http.NewServeMux()
	http.ListenAndServe(":8080", nil)
}
