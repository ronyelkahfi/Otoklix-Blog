package blogs

import (
	_blogDomain "otoklix-blog/business/blogs"
	_controllers "otoklix-blog/controllers"
	_response "otoklix-blog/controllers/blogs/response"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	usecase _blogDomain.Usecase
}

func NewBlogController(BlogUsecase _blogDomain.Usecase) *BlogController {
	return &BlogController{
		usecase: BlogUsecase,
	}
}
func (bc *BlogController) GetPost(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := bc.usecase.GetPosts(ctx)
	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.ListDomainToResponse(data))
}
