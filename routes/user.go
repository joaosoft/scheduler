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
		web.NewRoute(web.MethodGet, "/schedule/users/:id", controller.GetUser),
		web.NewRoute(web.MethodPost, "/schedule/users", controller.CreateUser),
		web.NewRoute(web.MethodPut, "/schedule/users/:id", controller.UpdateUser),
		web.NewRoute(web.MethodDelete, "/schedule/users/:id", controller.DeleteUser),
	)

	return err
}
