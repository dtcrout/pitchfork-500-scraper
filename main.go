package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) string {
	// Simple wrapper for http get() method
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	html := string(body)

	return html
}

func main() {
	url := "https://en.wikipedia.org/wiki/The_Pitchfork_500"

	html := get(url)

	fmt.Println(html)
}
