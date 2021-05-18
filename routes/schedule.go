package routes

import (
	"scheduler/controllers"
	"scheduler/models"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/logger"

	"github.com/joaosoft/web"
)

func RegisterScheduleRoutes(ns *web.Namespace, logger logger.ILogger, dbrConfig *dbr.DbrConfig) (err error) {
	model, err := models.NewScheduleModel(logger, dbrConfig)
	if err != nil {
		return err
	}
	controller := controllers.NewScheduleController(model)

	err = ns.AddRoutes(
		web.NewRoute(web.MethodPost, "/schedule", controller.CreateSchedule),
		web.NewRoute(web.MethodGet, "/schedule", controller.ListSchedule),
		web.NewRoute(web.MethodGet, "/schedule/:id", controller.GetSchedule),
		web.NewRoute(web.MethodPut, "/schedule/:id", controller.UpdateSchedule),
		web.NewRoute(web.MethodDelete, "/schedule/:id", controller.DeleteSchedule),
	)

	return err
}
