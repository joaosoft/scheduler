package routes

import (
	"scheduler/controllers"
	"scheduler/models"

	"github.com/joaosoft/dbr"

	"github.com/joaosoft/logger"

	"github.com/joaosoft/web"
)

func RegisterUserRoutes(ns *web.Namespace, logger logger.ILogger, dbrConfig *dbr.DbrConfig) (err error) {
	model, err := models.NewUserModel(logger, dbrConfig)
	if err != nil {
		return err
	}
	controller := controllers.NewUserController(model)

	err = ns.AddRoutes(
		web.NewRoute(web.MethodGet, "/schedule/user/:id", controller.GetUser),
		web.NewRoute(web.MethodPost, "/schedule/user", controller.CreateUser),
		web.NewRoute(web.MethodPut, "/schedule/user/:id", controller.UpdateUser),
	)

	return err
}
