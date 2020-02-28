package webscraping

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

func GetTittle(client *http.Client, link string) string {
	resp, err := client.Get(link)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if title, ok := getHtmlTitle(resp.Body); ok {
		log.Debug("Tittle website: ", title)
		return title
	} else {
		println("Fail to get HTML title")
	}

	return ""
}

func getHtmlTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}
