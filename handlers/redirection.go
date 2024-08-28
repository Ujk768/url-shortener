package handlers

import (
	"log"
	"net/http"

	"github.com/Ujk768/url-shortener/data"
)

type Redirection struct {
	l *log.Logger
}

func NewRedirection(l *log.Logger) *Redirection {
	return &Redirection{l}
}

// handle redirection of incoming requests
func (s *Redirection) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
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

func (s *Redirection) getRedirection(rw http.ResponseWriter, rq *http.Request) {

	s.l.Println("Inisde Get Redirection Handler")
	d := &data.GetRequestRedirection{}
	err := d.FromJSONRedirection(rq.Body)

	ru, err := data.GetRedirectionURL(d.ShortenUrl)
	if err != nil {
		http.Error(rw, "Unable to find URL", http.StatusNotFound)
	}

	http.Redirect(rw, rq, ru.LongURL, http.StatusTemporaryRedirect)

}
