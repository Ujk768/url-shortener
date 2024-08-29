package handlers

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/Ujk768/url-shortener/data"
)

type GetRequest struct {
	OriginalUrl string
	Domain      string
}

type GetResponse struct {
	Message    string
	ShortenUrl string
}

type Shortner struct {
	l *log.Logger
}

func NewShortner(l *log.Logger) *Shortner {
	return &Shortner{l}
}

// longurl -> shorturl
func ShortenURL(longurl string) string {
	// Hash the long URL using SHA-1
	h := sha1.New()
	h.Write([]byte(longurl))
	d := h.Sum(nil)

	// Encode the hash as a base64 string
	encodedString := base64.URLEncoding.EncodeToString(d)

	// Truncate the encoded string to the first 6 characters
	shortURL := encodedString[:6]
	return shortURL
}

func (s *Shortner) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	s.l.Println("Inside Shortner Handler")
	if rq.Method == http.MethodPost {
		s.AddUrl(rw, rq)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Shortner) AddUrl(rw http.ResponseWriter, rq *http.Request) {
	d := &data.GetRequestShortern{}
	err := d.FromJSON(rq.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal JSON", http.StatusBadRequest)
	}
	nu := data.AddURL(d.LongURL, d.Domain)
	err = nu.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
