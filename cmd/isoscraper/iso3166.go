package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/nklizhe/isoconv/iso3166"
)

func scrapeISO3166() ([]iso3166.Region, error) {
	// alpha-1
	doc, err := goquery.NewDocument("http://en.wikipedia.org/wiki/ISO_3166-1")
	if err != nil {
		return nil, err
	}

	all := []iso3166.Region{}
	doc.Find("#mw-content-text > table.wikitable:nth-of-type(1) > tbody > tr").Each(func(i int, s *goquery.Selection) {
		var r iso3166.Region
		r.Name = s.Find("td:nth-of-type(1) > a").Text()
		r.Alpha2 = s.Find("td:nth-of-type(2) > a > span").Text()
		r.Alpha3 = s.Find("td:nth-of-type(3) > span").Text()
		all = append(all, r)
	})

	// alpha3

	return all, nil
}
