package main

import (
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

func (t Track) String() string {
	return fmt.Sprintf("{\"artist\": \"%s\", \"song\": \"%s\"}", t.Artist, t.Song)
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
			track := Track{Artist: artist, Song: song}
			fmt.Println(track)
		})
	})
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"

	ParseWiki(url)
}
