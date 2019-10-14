package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func ParseWiki(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Find all songs on page and parse string into artist and song
	doc.Find(".div-col").Each(func(_ int, s *goquery.Selection) {
		s.Find("li").Each(func(_ int, t *goquery.Selection) {
			text := t.Text()
			s := strings.Split(text, " –")
			artist := s[0]
			song := s[1]
			fmt.Println(artist, song)
		})
	})
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"
	ParseWiki(url)
}
