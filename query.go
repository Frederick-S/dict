package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func query(name string) Word {
	url := "http://dict.youdao.com/w/" + name

	document, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	explanations := []string{}

	document.Find(".wordbook-js + .trans-container ul li").Each(func(i int, selection *goquery.Selection) {
		explanations = append(explanations, selection.Text())
	})

	word := Word{
		name:         name,
		explanations: explanations,
	}

	return word
}
