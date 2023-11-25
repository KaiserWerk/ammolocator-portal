package main

import (
	"regexp"
	"strconv"
	"strings"
)

var shutgunRegex1 = regexp.MustCompile("[0-9]{2}/[0-9]{2}")
var shutgunRegex2 = regexp.MustCompile("[0-9]{2}/[0-9]{2},[0-9]")
var shutgunRegex3 = regexp.MustCompile("[0-9]{2}/[0-9]{2}.[0-9]")

func detectBrand(input string) string {
	input = strings.ToLower(input)
	if strings.Contains(input, "s&b") || strings.Contains(input, "sellier&bellot") ||
		strings.Contains(input, "sellier & bellot") || strings.Contains(input, "s & b") {
		return "Sellier & Bellot"
	} else if strings.Contains(input, "geco") {
		return "Geco"
	} else if strings.Contains(input, "hornady") {
		return "Hornady"
	} else if strings.Contains(input, "prvi partizan") {
		return "Prvi Partizan"
	} else if strings.Contains(input, "aguila") {
		return "Aguila"
	} else if strings.Contains(input, "maxxtech") {
		return "MaxxTech"
	} else if strings.Contains(input, "speer") {
		return "Speer"
	} else if strings.Contains(input, "fiocchi") {
		return "Fiocchi"
	} else if strings.Contains(input, "topshot") {
		return "TopShot"
	} else if strings.Contains(input, "winchester") {
		return "Winchester"
	} else if strings.Contains(input, "magtech") {
		return "MagTech"
	} else if strings.Contains(input, "barnes") {
		return "Barnes"
	} else if strings.Contains(input, "ggg") {
		return "GGG"
	} else if strings.Contains(input, "stv") {
		return "STV"
	} else if strings.Contains(input, "barnaul") {
		return "Barnaul"
	} else if strings.Contains(input, "cci") {
		return "CCI"
	} else if strings.Contains(input, "blazer") {
		return "Blazer"
	} else if strings.Contains(input, "remington") {
		return "Remington"
	} else if strings.Contains(input, "federal premium") {
		return "Federal Premium"
	} else if strings.Contains(input, "eley") {
		return "Eley"
	} else if strings.Contains(input, "lapua") {
		return "Lapua"
	} else if strings.Contains(input, "norma") {
		return "Norma"
	} else if strings.Contains(input, "rws") {
		return "RWK"
	} else if strings.Contains(input, "sk") {
		return "SK"
	} else if strings.Contains(input, "bioammo") {
		return "BioAmmo"
	} else if strings.Contains(input, "rottweil") {
		return "Rottweil"
	} else if strings.Contains(input, "zink") {
		return "Zink"
	} else if strings.Contains(input, "brenneke") {
		return "Brenneke"
	} else if strings.Contains(input, "ddupleks") {
		return "DDupleks"
	} else if strings.Contains(input, "browning") {
		return "Browning"
	} else if strings.Contains(input, "nosler") {
		return "Nosler"
	} else if strings.Contains(input, "sako") {
		return "Sako"
	} else if strings.Contains(input, "swissp") {
		return "SwissP"
	} else if strings.Contains(input, "Weatherby") {
		return "Weatherby"
	}

	return "n/a"
}

func detectCaliber(input string) string {
	input = strings.ToLower(input)
	if strings.Contains(input, "22 lfb") || strings.Contains(input, "22 lr") ||
		strings.Contains(input, "22lfb") || strings.Contains(input, "22lr") {
		return ".22 lfb"
	} else if strings.Contains(input, "9mmluger") || strings.Contains(input, "9 mmluger") ||
		strings.Contains(input, "9mm luger") || strings.Contains(input, "9 mm luger") ||
		strings.Contains(input, "9x19") {
		return "9mm Luger"
	} else if strings.Contains(input, "45acp") || strings.Contains(input, "45 acp") ||
		strings.Contains(input, "45 auto") {
		return ".45 ACP"
	} else if strings.Contains(input, "38spl") || strings.Contains(input, "38 spl") ||
		strings.Contains(input, "38spec") || strings.Contains(input, "38 spec") ||
		strings.Contains(input, "38special") || strings.Contains(input, "38 special") {
		return ".38 Special"
	} else if strings.Contains(input, "357magnum") || strings.Contains(input, "357mag") ||
		strings.Contains(input, "357 magnum") || strings.Contains(input, "357 mag") {
		return ".357 Magnum"
	} else if strings.Contains(input, "7,62x39") || strings.Contains(input, "7.62x39") {
		return "7,62x39"
	} else if strings.Contains(input, "223rem") || strings.Contains(input, "223remington") ||
		strings.Contains(input, "223 rem") || strings.Contains(input, "223 remington") {
		return ".223 Remington"
	} else if strings.Contains(input, "308win") || strings.Contains(input, "308winchester") ||
		strings.Contains(input, "308 win") || strings.Contains(input, "308 winchester") {
		return ".308 Winchester"
	} else if strings.Contains(input, "6,5cr") || strings.Contains(input, "6,5creedmoor") ||
		strings.Contains(input, "6,5 cr") || strings.Contains(input, "6,5 creedmoor") ||
		strings.Contains(input, "6.5cr") || strings.Contains(input, "6.5creedmoor") ||
		strings.Contains(input, "6.5 cr") || strings.Contains(input, "6.5 creedmoor") ||
		strings.Contains(input, "6.5mmcr") || strings.Contains(input, "6.5mm creedmoor") ||
		strings.Contains(input, "6,5mmcr") || strings.Contains(input, "6,5mm creedmoor") {
		return "6,5 Creedmoor"
	} else if s := shutgunRegex1.FindString(input); s != "" {
		return s
	} else if s := shutgunRegex2.FindString(input); s != "" {
		return s
	} else if s := shutgunRegex3.FindString(input); s != "" {
		return s
	}

	return "n/a"
}

func detectPrice(input string) float64 {
	input = strings.Replace(strings.Trim(input, "€ "), ",", ".", 1)
	p, _ := strconv.ParseFloat(input, 64)
	return p
}
