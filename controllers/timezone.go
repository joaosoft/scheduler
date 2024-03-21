package controllers

import (
	"github.com/joaosoft/scheduler/models"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type TimezoneController struct {
	model *models.TimezoneModel
}

func NewTimezoneController(model *models.TimezoneModel) *TimezoneController {
	return &TimezoneController{
		model: model,
	}
}

func (c *TimezoneController) ListTimezone(ctx *web.Context) error {
	response, err := c.model.ListTimezone()
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *TimezoneController) GetTimezone(ctx *web.Context) error {
	request := &GetTimezoneRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.GetTimezone{
		Id: request.Id,
	}
	response, err := c.model.GetTimezone(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}
