package versions

import (
	"github.com/gofiber/fiber/v2"
)

func V1(app *fiber.App){

	v1 := app.Group("/v1")

	v1.Get("/", func (c *fiber.Ctx)error{
		return c.SendStatus(fiber.StatusOK)
	})

}