package routes

import (
	"scheduler/controllers"
	"scheduler/models"

	"github.com/joaosoft/dbr"

	"github.com/joaosoft/logger"

	"github.com/joaosoft/web"
)

func RegisterTimezoneRoutes(ns *web.Namespace, logger logger.ILogger, dbrConfig *dbr.DbrConfig) (err error) {
	model, err := models.NewTimezoneModel(logger, dbrConfig)
	if err != nil {
		return err
	}
	controller := controllers.NewTimezoneController(model)

	err = ns.AddRoutes(
		web.NewRoute(web.MethodGet, "/schedule/timezones", controller.ListTimezone),
		web.NewRoute(web.MethodGet, "/schedule/timezones/:id", controller.GetTimezone),
	)

	return err
}
