package url

import (
	hash2 "awesomeProject1/hash"
	"fmt"
)

const baseUrl = "https://me.li/%s"

func GenerateShortURL(originalUrl string) string {
	shortHash := enshortOriginalURL(originalUrl)

	return fmt.Sprintf(baseUrl, shortHash)
}

func enshortOriginalURL(originalURL string) string {
	hash := hash2.Encode(originalURL)
	return hash[:5]
}
