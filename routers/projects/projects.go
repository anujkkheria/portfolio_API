package projects

import (
	handler "github.com/anujkkheria/portfolio_API/handlers/projects"
	"github.com/gofiber/fiber/v2"
)

	func ProjDemoHandler(c *fiber.Ctx) error{
	return c.Status(200).SendString("you are in projects route")
	}

	func Projects(v *fiber.App){
		projects := v.Group("/projects")
		projects.Get("/",handler.GetProjects)
		projects.Post("/",handler.CreateNewProject)
		projects.Delete("/:id",handler.DeleteProject)
	}