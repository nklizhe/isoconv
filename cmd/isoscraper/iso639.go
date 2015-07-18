package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/nklizhe/isoconv/iso639"
)

func scrapeISO639() ([]iso639.Language, error) {
	// alpha-1
	doc, err := goquery.NewDocument("https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes")
	if err != nil {
		return nil, err
	}

	all := []iso639.Language{}
	doc.Find("#mw-content-text > table.wikitable > tbody > tr").Each(func(i int, s *goquery.Selection) {
		var r iso639.Language

		r.Family = s.Find("td:nth-of-type(2) > a").Text()
		r.Name = s.Find("td:nth-of-type(3) > a").Text()
		r.NativeName = s.Find("td:nth-of-type(4)").Text()
		r.Code1 = s.Find("td:nth-of-type(5)").Text()
		r.Code2 = s.Find("td:nth-of-type(6)").Text()
		r.Code2B = s.Find("td:nth-of-type(7)").Text()
		r.Code3 = s.Find("td:nth-of-type(8)").Text()
		r.Code6 = s.Find("td:nth-of-type(9)").Text()
		all = append(all, r)
	})

	// TODO: scrape alpha3

	return all, nil
}
