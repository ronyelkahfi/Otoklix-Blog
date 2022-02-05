package routes

import (
	_blogController "otoklix-blog/controllers/blogs"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	BlogController _blogController.BlogController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/posts", cl.BlogController.GetPost)
	c.GET("/posts/:id", cl.BlogController.ShowPost)
	c.POST("/posts", cl.BlogController.CreatePost)
	c.PUT("/posts/:id", cl.BlogController.UpdatePost)
	c.DELETE("/posts/:id", cl.BlogController.DeletePost)
}
