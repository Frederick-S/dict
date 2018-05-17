package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"

	"net/http"
	"strings"
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

	word := Word{
		name:         name,
		explanations: []string{},
		suggestions:  []string{},
	}

	document.Find(".wordbook-js + .trans-container ul li").Each(func(_ int, selection *goquery.Selection) {
		word.explanations = append(word.explanations, selection.Text())
	})

	if len(word.explanations) == 0 {
		document.Find(".typo-rel").Each(func(_ int, selection *goquery.Selection) {
			suggestion := strings.Replace(selection.Text(), " ", "", -1)
			suggestion = strings.Replace(suggestion, "\n", "", -1)

			word.suggestions = append(word.suggestions, suggestion)
		})
	}

	return word
}
