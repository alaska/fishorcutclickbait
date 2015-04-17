package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	cb "github.com/alaska/fishorcutclickbait"
	"os"
	"strings"
)

func main() {
	doc, err := goquery.NewDocument(`http://viralnova.com`)
	if err != nil {
		fmt.Printf("Attempt to pull from site failed: [%s]\n", err.Error())
		os.Exit(1)
	}
	total := 0
	clickbait := 0
	doc.Find(`h4.snip > a`).Each(func(i int, s *goquery.Selection) {
		total++
		title := strings.TrimSpace(s.Text())
		if cb.IsClickbait(title) {
			fmt.Println(title, "identified as clickbait")
			clickbait++
		}
	})
	fmt.Printf("%d/%d articles identified as clickbait\n", clickbait, total)
}
