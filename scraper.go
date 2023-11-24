package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func getFrankoniaProducts() []Product {
	const baseUrl = "https://www.frankonia.de"
	urls := []string{
		"https://www.frankonia.de/schiesssport/munition/kk-patronen/Artikel.html",
		"https://www.frankonia.de/schiesssport/munition/kurzwaffenpatronen/Artikel.html",
		"https://www.frankonia.de/schiesssport/munition/kurzwaffenpatronen/Artikel.html?page=1",
		"https://www.frankonia.de/schiesssport/munition/patronen-fuer-buechsen/Artikel.html",
		"https://www.frankonia.de/schiesssport/munition/patronen-fuer-buechsen/Artikel.html?page=1",
		"https://www.frankonia.de/schiesssport/munition/patronen-fuer-flinten/Artikel.html",
	}

	products := make([]Product, 0)

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("div.fr-article-tile.fr-article-tile--square", func(e *colly.HTMLElement) {
		p := Product{}

		p.URL = baseUrl + e.ChildAttr("a.fr-article-tile__image-link", "href")
		p.Image = e.ChildAttr("img.fr-article-tile__image", "src")
		p.Retailer = "Frankonia"
		p.Brand = e.ChildText("div.fr-article-tile__brand")
		p.ProductName = strings.TrimSpace(e.ChildText("div.fr-article-tile__name"))
		p.Caliber = detectCaliber(p.ProductName)
		price := strings.TrimSpace(e.ChildText("ins.fr-price--ins"))
		if price == "" {
			price = e.ChildText("span.fr-price-min")
		}
		//fmt.Println("name:", p.ProductName, "price:", price)
		p.Price = detectPrice(price)

		products = append(products, p)
	})

	c.OnError(func(resp *colly.Response, err error) {
		fmt.Println(err)
	})

	for _, v := range urls {
		c.Visit(v)
	}

	return products
}
