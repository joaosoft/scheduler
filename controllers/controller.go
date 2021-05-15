package controllers

import (
	"scheduler/models"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type Controller struct {
	model *models.Model
}

func NewController(model *models.Model) *Controller {
	return &Controller{
		model: model,
	}
}

func (c *Controller) ListTimezone(ctx *web.Context) error {
	response, err := c.model.ListTimezone()
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) GetTimezone(ctx *web.Context) error {
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

func (c *Controller) ListSchedule(ctx *web.Context) error {
	response, err := c.model.ListSchedule()
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) GetSchedule(ctx *web.Context) error {
	request := &GetScheduleRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.GetSchedule{
		Id: request.Id,
	}
	response, err := c.model.GetSchedule(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) CreateSchedule(ctx *web.Context) error {
	request := &CreateScheduleRequest{}

	if err := ctx.Request.Bind(&request.Body); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.CreateSchedule{
		Subject:     request.Body.Subject,
		Description: request.Body.Description,
		IdTimezone:  request.Body.IdTimezone,
	}
	response, err := c.model.CreateSchedule(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusCreated, response)
}

func (c *Controller) UpdateSchedule(ctx *web.Context) error {
	request := &UpdateScheduleRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.UpdateSchedule{
		Subject:     request.Body.Subject,
		Description: request.Body.Description,
		IdTimezone:  request.Body.IdTimezone,
	}
	response, err := c.model.UpdateSchedule(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) DeleteSchedule(ctx *web.Context) error {
	request := &DeleteScheduleRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	param := &models.DeleteSchedule{
		Id: request.Id,
	}
	err := c.model.DeleteSchedule(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.NoContent(web.StatusOK)
}
