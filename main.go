package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Product struct {
		URL         string
		Image       string
		Retailer    string
		Brand       string
		ProductName string
		Caliber     string
		Price       float64
	}
)

func main() {
	allProducts := make([]Product, 0, 500)

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	allProducts = append(allProducts, getFrankoniaProducts()...)

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	headers := []string{
		"URL",
		//"Image",
		"Retailer",
		"Brand",
		"Product Name",
		"Caliber",
		"Price",
	}
	writer.Write(headers)

	for _, v := range allProducts {
		if v.Caliber == "n/a" {
			continue
		}

		record := []string{
			v.URL,
			//v.Image,
			v.Retailer,
			v.Brand,
			v.ProductName,
			v.Caliber,
			strings.Replace(strconv.FormatFloat(v.Price, 'f', 2, 64), ".", ",", 1),
		}

		writer.Write(record)
	}
	writer.Flush()
}
