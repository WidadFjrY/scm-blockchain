package helper

import (
	"fmt"
	"strings"
)

func FormatRupiah(number float32) string {
	// Ubah float menjadi string dengan format 2 desimal
	strNum := fmt.Sprintf("%.2f", number)

	// Pisahkan bagian sebelum dan setelah koma
	parts := strings.Split(strNum, ".")
	intPart := parts[0]

	n := len(intPart)
	if n <= 3 {
		return "Rp. " + intPart
	}

	// Menyimpan hasil dengan titik sebagai pemisah ribuan
	result := ""
	counter := 0

	// Iterasi dari belakang ke depan
	for i := n - 1; i >= 0; i-- {
		result = string(intPart[i]) + result
		counter++
		if counter%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp. " + result
}
