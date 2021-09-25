package main

import (
	search "Beep/search"
	"fmt"
)

func main() {
	res, _ := search.LoadDocuments("wiki.xml")

	xs := search.Search(res, "cat")

	for _, docUrl := range xs {
		fmt.Println(docUrl.URL)
	}
}
