package controllers

import (
	"scheduler/models"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(model *models.UserModel) *UserController {
	return &UserController{
		model: model,
	}
}

func (c *UserController) GetUser(ctx *web.Context) error {
	request := &GetUserRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.GetUser{
		Id: request.Id,
	}
	response, err := c.model.GetUser(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *UserController) CreateUser(ctx *web.Context) error {
	request := &CreateUserRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if err := ctx.Request.Bind(&request.Body); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.CreateUser{
		Id:         request.Body.Id,
		IdCountry:  request.Body.IdCountry,
		IdTimezone: request.Body.IdTimezone,
	}
	response, err := c.model.CreateUser(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusCreated, response)
}

func (c *UserController) UpdateUser(ctx *web.Context) error {
	request := &UpdateUserRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.UpdateUser{
		Id:         request.Body.Id,
		IdCountry:  request.Body.IdCountry,
		IdTimezone: request.Body.IdTimezone,
	}
	response, err := c.model.UpdateUser(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *UserController) DeleteUser(ctx *web.Context) error {
	request := &DeleteUserRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.DeleteUser{
		Id: request.Id,
	}
	err := c.model.DeleteUser(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.NoContent(web.StatusOK)
}
