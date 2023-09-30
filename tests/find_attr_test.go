package tests

import (
	"io"
	"puzzle-english-translator/internal"
	"strings"
	"testing"
)

func TestFindAttr(t *testing.T) {

	htmlContent := `<html>
		<head>
			<meta property="og:description" content="Test Description"/>
			<meta property="og:image" content="https://example.com/image.jpg"/>
		</head>
	</html>`

	testReadCloser := io.NopCloser(strings.NewReader(htmlContent))

	puzzle := internal.FindAttr(testReadCloser)

	expectedText := "Test Description"
	expectedImgLink := "https://example.com/image.jpg"

	if puzzle.Text != expectedText {
		t.Errorf("Expected text: %s, got: %s", expectedText, puzzle.Text)
	}

	if puzzle.ImgLink != expectedImgLink {
		t.Errorf("Expected imgLink: %s, got: %s", expectedImgLink, puzzle.ImgLink)
	}
}
