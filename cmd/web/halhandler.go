package web

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HalHandler(c *fiber.Ctx) error {
	message := "I’m sorry, Dave. I’m afraid I can’t do that."
	component := HalResponse(message)

	// The adaptor makes the templ.Handler compatible with Fiber's handler signature.
	// It automatically sets the Content-Type to "text/html".
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}
