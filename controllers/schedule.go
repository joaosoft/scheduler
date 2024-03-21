package controllers

import (
	"github.com/joaosoft/scheduler/models"

	uuid "github.com/satori/go.uuid"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type ScheduleController struct {
	model *models.ScheduleModel
}

func NewScheduleController(model *models.ScheduleModel) *ScheduleController {
	return &ScheduleController{
		model: model,
	}
}

func (c *ScheduleController) ListSchedule(ctx *web.Context) error {
	response, err := c.model.ListSchedule()
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *ScheduleController) GetSchedule(ctx *web.Context) error {
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

func (c *ScheduleController) CreateSchedule(ctx *web.Context) error {
	request := &CreateScheduleRequest{}

	if err := ctx.Request.BindUrlParams(&request); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if err := ctx.Request.Bind(&request.Body); err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	uid := uuid.NewV4()

	param := &models.CreateSchedule{
		HashedId:         uid.String(),
		Subject:          request.Body.Subject,
		Description:      request.Body.Description,
		IdUser:           request.Body.IdUser,
		IdTimezone:       request.Body.IdTimezone,
		TimeSlots:        request.Body.TimeSlots,
		IdScheduleStatus: models.ScheduleStatusNewID,
	}
	response, err := c.model.CreateSchedule(param)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, err)
	}

	return ctx.Response.JSON(web.StatusCreated, response)
}

func (c *ScheduleController) UpdateSchedule(ctx *web.Context) error {
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

func (c *ScheduleController) DeleteSchedule(ctx *web.Context) error {
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
