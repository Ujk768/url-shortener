package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type UrlMapping struct {
	OriginalUrl string
}

type GetRequestShortern struct {
	longURL string
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
}

var urlList = make(map[string]UrlMapping)

var ErrorNotFound = fmt.Errorf("URL Not Found")

// map[short_url]= long_url
func AddURL(longURL string, shortURL string) {
	urlList[shortURL] = UrlMapping{
		OriginalUrl: longURL,
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
