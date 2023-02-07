package main

import (
	"github.com/PuerkitoBio/goquery"
)

func hangler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}
}
