package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"time"
)

const URL = "https://puzzle-english.com/fun/"

type PuzzleEnglishDailyTranslation struct {
	text    string
	imgLink string
}

func main() {

	translate := getTranslateRequest()
	findAttr := findAttr(translate)

	fmt.Println(findAttr.text, findAttr.imgLink)

	defer translate.Close()

}

func findAttr(readCloser io.ReadCloser) PuzzleEnglishDailyTranslation {

	doc, err := goquery.NewDocumentFromReader(readCloser)
	puzzle := PuzzleEnglishDailyTranslation{}
	if err != nil {
		panic(err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {

		property, _ := s.Attr("property")
		content, _ := s.Attr("content")

		if property == "og:description" {
			puzzle.text = content
		}

		if property == "og:image" {
			puzzle.imgLink = content
		}
	})

	return puzzle
}

func getTranslateRequest() io.ReadCloser {

	res, err := http.Get(URL + time.Now().Format("02-01-2006"))
	if err != nil {
		panic(err)
	}

	return res.Body
}
