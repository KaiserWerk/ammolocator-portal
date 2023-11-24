package main

import (
	"strconv"
	"strings"
)

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
		strings.Contains(input, "6.5 cr") || strings.Contains(input, "6.5 creedmoor") {
		return "6,5 Creedmoor"
	}

	return "n/a"
}

func detectPrice(input string) float64 {
	input = strings.Replace(strings.Trim(input, "€ "), ",", ".", 1)
	p, _ := strconv.ParseFloat(input, 64)
	return p
}
