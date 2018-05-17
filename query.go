package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func query(name string) Word {
	url := "http://dict.youdao.com/w/" + name
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var explanations []string = nil

	document.Find(".wordbook-js + .trans-container ul li").Each(func(i int, selection *goquery.Selection) {
		explanations = append(explanations, selection.Text())
	})

	word := Word{
		name:         name,
		explanations: explanations,
	}

	return word
}
