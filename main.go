package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func ParseWiki(url string) []string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Create array for tracks
	var tracks []string
	var track string

	tracks = append(tracks, "artist,song,\n")

	// Find all songs on page and parse string into artist and song
	doc.Find(".div-col").Each(func(_ int, s *goquery.Selection) {
		s.Find("li").Each(func(_ int, t *goquery.Selection) {
			text := strings.Split(t.Text(), " –")

			artist := text[0]
			song := strings.Trim(text[1], " \"")

			// Create track
			track = artist + "," + song + "\n"

			tracks = append(tracks, string(track))
		})
	})

	return tracks
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"
	out_filename := "tracks.csv"

	tracks := ParseWiki(url)

	file, _ := os.Create(out_filename)

	defer file.Close()

	var err error
	for _, track := range tracks {
		_, err = io.WriteString(file, track)

		if err != nil {
			log.Fatal(err)
		}

		file.Sync()
	}
}
