package utils

import "github.com/skip2/go-qrcode"

func GenerateASCII(text string) string {

	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return "Error generating QR"
	}

	bitmap := qr.Bitmap()

	ascii := ""

	for _, row := range bitmap {
		for _, col := range row {
			if col {
				ascii += "██"
			} else {
				ascii += "  "
			}
		}
		ascii += "\n"
	}

	return ascii
}