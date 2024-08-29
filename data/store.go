package data

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type UrlMapping struct {
	OriginalUrl string
	Domain      string
}

type GetRequestShortern struct {
	LongURL string `json:"longUrl"`
	Domain  string `json:"domain"`
}

type GetResponseShorten struct {
	Message  string `json:"message"`
	ShortUrl string `json:"shortUrl"`
}
type GetRequestRedirection struct {
	ShortenUrl string `json:"shortUrl"`
}

type GetResponseRedirection struct {
	LongURL string `json:"longUrl"`
	Domain  string `json:"domain"`
}

var urlList = make(map[string]UrlMapping)

var ErrorNotFound = fmt.Errorf("URL Not Found")

// map[short_url]= long_url
func AddURL(longURL string, domain string) GetResponseShorten {
	var su string
	var exists bool

	// Try generating a unique short URL
	for i := 0; i < 5; i++ { // limiting to 5 attempts to avoid an infinite loop
		su = ShortenURL(longURL + string(i)) // Modify the input slightly if necessary
		_, exists = urlList[su]
		if !exists {
			break // If the short URL doesn't exist, we found a unique one
		}
	}

	if exists {
		// Handle the case where all attempts to find a unique URL failed
		return GetResponseShorten{
			ShortUrl: "",
			Message:  "Failed to create a unique short URL after multiple attempts",
		}
	}

	// Store the unique short URL
	urlList[su] = UrlMapping{
		OriginalUrl: longURL,
		Domain:      domain,
	}

	return GetResponseShorten{
		ShortUrl: su,
		Message:  "Created new short URL",
	}
}

func GetMap() map[string]UrlMapping {
	return urlList
}

func GetRedirectionURL(url string) (GetResponseRedirection, error) {
	longurl, ok := urlList[url]

	if !ok {
		return GetResponseRedirection{}, ErrorNotFound
	}

	return GetResponseRedirection{
		LongURL: longurl.OriginalUrl,
		Domain:  longurl.Domain,
	}, nil

}

func DeleteUrl(url string) {
	delete(urlList, url)
}

func (res *GetResponseShorten) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(res)
}

func (res *GetResponseRedirection) ToJSONRedirection(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(res)
}

func (u *GetRequestShortern) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (g *GetRequestRedirection) FromJSONRedirection(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(g)
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

func CleanDomain(domain string) string {
	parsedUrl, err := url.Parse(domain)
	if err != nil {
		return domain // return as-is if parsing fails
	}

	// If the scheme exists, strip it
	if parsedUrl.Scheme != "" {
		return strings.TrimPrefix(domain, parsedUrl.Scheme+"://")
	}
	return domain
}
