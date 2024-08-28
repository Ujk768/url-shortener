package data

import (
	"encoding/json"
	"io"
)

type UrlMapping struct {
	Domain      string
	OriginalUrl string
	ShortUrl    string
}

type GetResponse struct {
	Message  string
	ShortUrl string
}
type GetRequest struct {
	ShortUrl string
}

var urlList = make(map[string]UrlMapping)

// map[short_url]= long_url
func AddURL(domain string, originalURL string, shortURL string) {
	urlList[shortURL] = UrlMapping{
		Domain:      domain,
		OriginalUrl: originalURL,
		ShortUrl:    shortURL,
	}
}

func GetMap() map[string]UrlMapping {
	return urlList
}

// func SearchOriginalUrl(url string) string {
// 	value := urlList[url]
// 	return value.OriginalUrl
// }

func GetRedirectionURL(url string) GetResponse {
	value := urlList[url]
	if value.ShortUrl == "" {
		res := GetResponse{
			Message: "Redirection URL not found",
		}
		return res
	}
	res := GetResponse{
		Message:  "Got Redirection url",
		ShortUrl: value.ShortUrl,
	}
	return res

}

func DeleteUrl(url string) {
	delete(urlList, url)
}

func (res *GetResponse) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(res)
}

func (u *GetRequest) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
