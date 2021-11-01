package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://countrycode.org/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var prefixs []string

	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			s.Find("tr").Each(func(idxtr int, tr *goquery.Selection) {
				tr.Find("td").Each(func(idxtd int, td *goquery.Selection) {
					if idxtd == 1 {
						pfx := strings.Split(td.Text(), ",")
						for _, i := range pfx {
							i = strings.ReplaceAll(i, "-", "")
							i = strings.ReplaceAll(i, " ", "")
							prefixs = append(prefixs, fmt.Sprintf("+%s", i))
						}
					}
				})
			})
		}
	})

	fmt.Println(strings.Join(prefixs, ","))
}
