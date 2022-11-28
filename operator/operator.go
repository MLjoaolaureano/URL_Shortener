package operator

import (
	database "awesomeProject1/storage"
	urlEnshorter "awesomeProject1/url"
	"strings"
)

func RecoverOriginalURL(url string) string {
	originalURL := database.RecoverReference(url)
	return originalURL
}
func CreateNewShortURL(url string) string {
	shortUrl := urlEnshorter.GenerateShortURL(url)
	lastFrontSlashIdx := strings.LastIndex(shortUrl, `/`)
	code := shortUrl[lastFrontSlashIdx+1:]

	database.CreateReference(code, url)
	return shortUrl
}
