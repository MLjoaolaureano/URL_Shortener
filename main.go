package main

import (
	"awesomeProject1/operator"
	"fmt"
)

func main() {
	url := "https://www.mercadolivre.com.br/"
	//wrongUrl := "https://www.marcadolibre.com.br/"
	//myHash := hash.Encode(url)
	//myWrongHash := hash.Encode(wrongUrl)
	//isSameHash := hash.Compare(url, myHash)
	//isWrongHash := hash.Compare(url, myWrongHash)

	//fmt.Printf("Current Hash : %s\n", myHash)
	//fmt.Printf("URL %s has same Hash : %t\n", url, isSameHash)
	//fmt.Printf("URL %s has same Hash : %t\n", wrongUrl, isWrongHash)

	shortUrl := operator.CreateNewShortURL(url)
	fmt.Printf("Current URL : %s\n", url)
	fmt.Printf("New URL : %s\n", shortUrl)

	recoveredUrl := operator.RecoverOriginalURL(shortUrl)

	fmt.Printf("Short URL : %s\n", shortUrl)
	fmt.Printf("Recovered URL : %s\n", recoveredUrl)
}
