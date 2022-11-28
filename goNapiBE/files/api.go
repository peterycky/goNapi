package files

import (
	"github.com/gofiber/fiber/v2"
)

var root string = "/Users/peterycky/Documents/"

func GetApi(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func GetDirectory(c *fiber.Ctx) error {
	directories := GetFullDirectoryList(root)
	return c.JSON(directories)
}

func GetFiles(c *fiber.Ctx) error {

	extensions := []string{".pdf", ".txt", ".js"}

	param := c.Query("dir")

	directory := root + param
	directories := append(ObtainDirectoryList(root), param)

	if !ContainsString(directories, param) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	files := ObtainFileList(directory, extensions)
	return c.JSON(files)
}
