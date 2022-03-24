package scrapers

import (
	"strconv"
	"strings"
)

func ConvertePreco(preco string) float64 {

	precoFinal := strings.ReplaceAll(strings.TrimSpace(preco), "€", "")
	precoFinal = strings.ReplaceAll(strings.TrimSpace(precoFinal), " ", "")

	if s, err := strconv.ParseFloat(precoFinal, 64); err == nil {
		return s
	}

	return 0
}
