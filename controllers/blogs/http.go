package blogs

import (
	_blogDomain "otoklix-blog/business/blogs"
	_controllers "otoklix-blog/controllers"
	_request "otoklix-blog/controllers/blogs/request"
	_response "otoklix-blog/controllers/blogs/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	usecase _blogDomain.Usecase
}

func NewController(BlogUsecase _blogDomain.Usecase) *BlogController {
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

func (bc *BlogController) ShowPost(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := bc.usecase.ShowPost(ctx, uint(id))
	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.DomainToResponse(data))
}

func (bc *BlogController) CreatePost(c echo.Context) error {
	ctx := c.Request().Context()
	var data _request.BlogRequest
	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	dataDomain := _request.RequestToDomain(data)
	dataResponse, err := bc.usecase.CreatePost(ctx, dataDomain)
	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.DomainToResponse(dataResponse))
}

func (bc *BlogController) UpdatePost(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	var data _request.BlogRequest
	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	dataDomain := _request.RequestToDomain(data)
	dataDomain.Id = uint(id)

	dataResponse, err := bc.usecase.UpdatePost(ctx, dataDomain)
	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.DomainToResponse(dataResponse))
}
func (bc *BlogController) DeletePost(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	dataResponse, err := bc.usecase.DeletePost(ctx, uint(id))
	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.DomainToResponse(dataResponse))
}
