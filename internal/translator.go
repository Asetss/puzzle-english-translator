package internal

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

type PuzzleEnglishDailyTranslation struct {
	Text    string
	ImgLink string
}

func FindAttr(readCloser io.ReadCloser) PuzzleEnglishDailyTranslation {

	doc, err := goquery.NewDocumentFromReader(readCloser)
	puzzle := PuzzleEnglishDailyTranslation{}
	if err != nil {
		panic(err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {

		property, _ := s.Attr("property")
		content, _ := s.Attr("content")

		if property == "og:description" {
			puzzle.Text = content
		}

		if property == "og:image" {
			puzzle.ImgLink = content
		}
	})

	return puzzle
}
