package handler

import (
	"awesomeProject1/operator"
	"fmt"
	"github.com/gofiber/fiber"
)

func CreateShortUrl(c *fiber.Ctx) {
	type Payload struct {
		Url string `json:"url"`
	}

	payload := new(Payload)

	if err := c.BodyParser(payload); err != nil {
		r := new(HttpResponse)
		r.success = true
		r.message = err.Error()
		r.httpStatusCode = 500
		response(r, c)
		return
	}

	shortUrl := operator.CreateNewShortURL(payload.Url)

	r := new(HttpResponse)
	r.success = true
	r.message = fmt.Sprintf("Short URL created : %s", shortUrl)
	r.httpStatusCode = 200
	response(r, c)
	return

}

func RecoverUrl(c *fiber.Ctx) {

	payload := c.Params("url")

	if payload == "" {
		r := new(HttpResponse)
		r.success = true
		r.message = "URL is empty"
		r.httpStatusCode = 500
		response(r, c)
		return
	}

	originalUrl := operator.RecoverOriginalURL(payload)

	r := new(HttpResponse)
	r.success = true
	r.message = fmt.Sprintf("URL recovered : %s", originalUrl)
	r.httpStatusCode = 200
	response(r, c)
	return
}

type HttpResponse struct {
	success        bool
	message        string
	httpStatusCode int
}

func response(httpResponse *HttpResponse, c *fiber.Ctx) {

	c.Status(httpResponse.httpStatusCode).JSON(&fiber.Map{
		"success": httpResponse.success,
		"message": httpResponse.message,
	})
	return
}
