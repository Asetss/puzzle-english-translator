package main

import (
	"fmt"
	"puzzle-english-translator/internal"
)

func main() {

	translate := internal.GetTranslateRequest()
	findAttr := internal.FindAttr(translate)

	fmt.Println(findAttr.Text, findAttr.ImgLink)

	defer translate.Close()

}
