package blogs

import (
	blogsHandler "github.com/anujkkheria/portfolio_API/handlers/blogs"
	"github.com/gofiber/fiber/v2"
)

func DemoHandler(c *fiber.Ctx) error{
	return c.Status(200).SendString("you visited blogs")
}

func Blogs(v *fiber.App){
	blogs := v.Group("/blogs")
	blogs.Get("/", DemoHandler)
	blogs.Get("/getall", blogsHandler.GetBlogs)
	blogs.Post("/createBlog", blogsHandler.CreateBlogs)
	blogs.Patch("/:id", blogsHandler.UpdateBlog)
}

