package utils

import (
	"fmt"
	"strings"
)

func FormatRupiah(amount float64) string {
	amountStr := fmt.Sprintf("%.2f", amount)

	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	n := len(integerPart)
	var formattedInt string
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			formattedInt += ","
		}
		formattedInt += string(integerPart[i])
	}

	return fmt.Sprintf("Rp %s.%s", formattedInt, decimalPart)
}

func BoolToYesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}
