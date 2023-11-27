package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
)

func getSSZBerkaProducts() []Product {
	products := make([]Product, 0)
	const (
		baseURL      = "https://www.schiesssportzentrum-berka.de/munition/"
		retailer     = "SSZ Berka"
		smallBoreCal = ".22 lfb"
	)
	urls := []string{
		"https://www.schiesssportzentrum-berka.de/munition/kleinkalibermunition/",
		"https://www.schiesssportzentrum-berka.de/munition/kurzwaffenmunition/",
		"https://www.schiesssportzentrum-berka.de/munition/langwaffenmunition/",
		"https://www.schiesssportzentrum-berka.de/munition/flinte-schrotpatronen/",
	}

	c := colly.NewCollector()
	c.UserAgent = userAgent

	c.OnHTML("div.module-type-text.diyfeLiveArea > table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, elem *colly.HTMLElement) {
			p := Product{}
			p.URL = baseURL
			name := elem.ChildText("td:nth-child(1)")
			p.ProductName = name
			packSize := strings.TrimSpace(elem.ChildText("td:nth-child(3)"))
			if packSize != "" {
				p.ProductName += " " + packSize + " Stk."
			}
			p.Retailer = retailer
			p.Brand = detectBrand(p.ProductName)
			p.Caliber = detectCaliber(p.ProductName)
			if p.Caliber == "n/a" {
				p.Caliber = smallBoreCal
			}
			p.Price = detectPrice(elem.ChildText("td:nth-child(4)"))
			products = append(products, p)

			price2 := strings.TrimSpace(elem.ChildText("td:nth-child(5)"))
			if price2 != "" && strings.HasSuffix(name, "€") {
				p2 := Product{}
				p2.URL = baseURL
				p2.Retailer = retailer
				p2.ProductName = name + " 500 Stk."
				p2.Price = detectPrice(price2)
				p2.Brand = detectBrand(name)
				p2.Caliber = detectCaliber(name)
				if p2.Caliber == "n/a" {
					p2.Caliber = smallBoreCal
				}
			}

			price3 := strings.TrimSpace(elem.ChildText("td:nth-child(6)"))
			if price3 != "" && strings.HasSuffix(name, "€") {
				p3 := Product{}
				p3.URL = baseURL
				p3.Retailer = retailer
				p3.ProductName = name + " 1000 Stk."
				p3.Price = detectPrice(price3)
				p3.Brand = detectBrand(name)
				p3.Caliber = detectCaliber(name)
				if p3.Caliber == "n/a" {
					p3.Caliber = smallBoreCal
				}
			}

			price4 := strings.TrimSpace(elem.ChildText("td:nth-child(7)"))
			if price4 != "" && strings.HasSuffix(name, "€") {
				p4 := Product{}
				p4.URL = baseURL
				p4.Retailer = retailer
				p4.ProductName = name + " 5000 Stk."
				p4.Price = detectPrice(price4)
				p4.Brand = detectBrand(name)
				p4.Caliber = detectCaliber(name)
				if p4.Caliber == "n/a" {
					p4.Caliber = smallBoreCal
				}
			}

		})

	})

	for _, v := range urls {
		c.Visit(v)
	}

	return products
}

func getArms24Products() []Product {
	products := make([]Product, 0)

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	urls := []string{
		"https://www.arms24.com/kleinkaliber-munition",
		"https://www.arms24.com/revolver-munition",
		"https://www.arms24.com/pistolen-munition",
		"https://www.arms24.com/schrotpatronen",
		"https://www.arms24.com/flintenlaufgeschosse",
		"https://www.arms24.com/buechsen-munition",
	}

	for _, u := range urls {
		var nodes []*cdp.Node
		chromedp.Run(ctx,
			chromedp.Navigate(u),
			chromedp.Nodes("div.product--box.box--minimal", &nodes, chromedp.ByQueryAll),
		)

		var url, name, price string
		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.AttributeValue("a.product--image", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("a.product--title.is--standard", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("span.price--default.is--nowrap", &price, chromedp.ByQuery, chromedp.FromNode(node)),
			)

			p := Product{}
			p.URL = url
			p.Image = ""
			p.Retailer = "Arms24"
			p.Brand = detectBrand(name)
			p.ProductName = strings.TrimSpace(name)
			p.Caliber = detectCaliber(name)
			p.Price = detectPrice(price)

			products = append(products, p)
		}
	}

	return products
}

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
	c.UserAgent = userAgent

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
