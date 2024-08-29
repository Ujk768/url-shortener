package handlers

import (
	"fmt"
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
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (s *Redirection) getRedirection(rw http.ResponseWriter, rq *http.Request) {

	s.l.Println("Inisde Get Redirection Handler")
	// d := &data.GetRequestRedirection{}
	// err := d.FromJSONRedirection(rq.Body)
	shortUrl := rq.URL.Path
	shortUrl = shortUrl[1:]
	s.l.Println("URL PATH", shortUrl)

	ru, err := data.GetRedirectionURL(shortUrl)
	if err != nil {
		http.Error(rw, "Unable to find URL", http.StatusNotFound)
	}
	sd := data.CleanDomain(ru.Domain)
	s.l.Println("sanitized DOMAIN", sd)
	redirectUrl := "https://" + sd + ru.LongURL
	fmt.Println("RedirectionURL", redirectUrl)
	http.Redirect(rw, rq, redirectUrl, http.StatusTemporaryRedirect)

}
