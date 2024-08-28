package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ujk768/url-shortener/data"
)

type Shortener struct {
	l *log.Logger
}

func (s *Shortener) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	// handle redirection
	if rq.Method == http.MethodGet {
		s.getRedirection(rw, rq)
		return
	}
	// Modify to list
	if rq.Method == http.MethodPut {

	}
	//add a url
	if rq.Method == http.MethodPost {

	}
	//delete from list
	if rq.Method == http.MethodDelete {

	}

}

func (s *Shortener) getRedirection(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println("Indide get Handler")
	d := data.GetRequest{}
	err := d.FromJSON(rq.Body)
	if err != nil {
		http.Error(rw, "Error in Get Redirection", http.StatusInternalServerError)
	}
	data.GetRedirectionURL(d.ShortUrl)

}
