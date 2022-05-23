package routes

import (
	"github.com/Jonathan-Bello/Api-portafolio/handlers/author"
	"github.com/Jonathan-Bello/Api-portafolio/handlers/blog"
	"github.com/Jonathan-Bello/Api-portafolio/handlers/tech"
	"github.com/labstack/echo/v4"
)

// Router declares the routes for the API
func Router(e echo.Echo) {
	e.GET("", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	AuthorRoutes(e)
	TechRoutes(e)
	BlogRoutes(e)
}

// AuthorRoutes declares the routes for the author
func AuthorRoutes(e echo.Echo) {
	authors := e.Group("/authors")
	authors.GET("", author.GetAll)
	authors.GET("/:id", author.GetByID)
	authors.POST("", author.Create)
	authors.PUT("", author.Update)
	authors.DELETE("/:id", author.Delete)
}

// TechRoutes declares the routes for the tech
func TechRoutes(e echo.Echo) {
	techs := e.Group("/techs")
	techs.GET("", tech.GetAll)
	techs.GET("/:id", tech.GetByID)
	techs.POST("", tech.Create)
	techs.PUT("", tech.Update)
	techs.DELETE("/:id", tech.Delete)
}

// BlogRoutes declares the routes for the blog
func BlogRoutes(e echo.Echo) {
	blogs := e.Group("/blogs")
	blogs.GET("", blog.GetAll)
	blogs.GET("/:id", blog.GetByID)
	blogs.POST("", blog.Create)
	blogs.PUT("", blog.Update)
	blogs.DELETE("/:id", blog.Delete)
}
