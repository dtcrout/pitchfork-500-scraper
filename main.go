package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Track struct {
	Artist string
	Song   string
}

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
			text := strings.Split(t.Text(), " –")

			artist := text[0]
			song := strings.Trim(text[1], " \"")

			// Create track
			track_obj := Track{Artist: artist, Song: song}
			track, _ := json.Marshal(track_obj)

			fmt.Println(string(track))
		})
	})
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"
	ParseWiki(url)
}
