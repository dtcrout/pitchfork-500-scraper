package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = io.WriteString(file, data)

	if err != nil {
		log.Fatal(err)
	}

	return file.Sync()
}

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

	tracks = append(tracks, "artist,song")

	// Find all songs on page and parse string into artist and song
	doc.Find(".div-col").Each(func(_ int, s *goquery.Selection) {
		s.Find("li").Each(func(_ int, t *goquery.Selection) {
			text := strings.Split(t.Text(), " –")

			artist := text[0]
			song := strings.Trim(text[1], " \"")

			// Create track
			track = artist + "," + song

			tracks = append(tracks, string(track))
		})
	})

	return tracks
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"
	out_filename := "tracks.csv"

	tracks := ParseWiki(url)

	var err error
	for _, track := range tracks {
		err = WriteToFile(out_filename, track)

		if err != nil {
			log.Fatal(err)
		}
	}
}
