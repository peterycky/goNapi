package files

import "github.com/gofiber/fiber/v2"

func ExecQnapi(c *fiber.Ctx) error {
	c.Response().SetStatusCode(200)
	c.SendString("qnapi executed")
	return nil
}
