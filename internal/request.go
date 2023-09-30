package internal

import (
	"io"
	"net/http"
	"time"
)

const URL = "https://puzzle-english.com/fun/"

func GetTranslateRequest() io.ReadCloser {

	res, err := http.Get(URL + time.Now().Format("02-01-2006"))
	if err != nil {
		panic(err)
	}

	return res.Body
}
