package analysis

import (
	"example.com/module/typefile"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
)

func Analize(u string) (Url typefile.UrlStruct) {

	baseUrl, err := url.Parse(u)
	if err != nil {
		return
	}

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	urls := Fetch(baseUrl, doc)
	isWp := CheckWp(baseUrl, doc)

	Url = typefile.UrlStruct{urls, err, isWp}

	return
}
