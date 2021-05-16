package routes

import (
	"scheduler/controllers"

	"github.com/joaosoft/web"
)

func RegisterRoutes(w *web.Server, c *controllers.Controller) (err error) {
	err = w.AddNamespace("api/v1/schedule").AddRoutes(
		// timezones
		web.NewRoute(web.MethodGet, "/schedule/timezones", c.ListTimezone),
		web.NewRoute(web.MethodGet, "/schedule/timezones/:id", c.GetTimezone),
	)

	if err != nil {
		return err
	}

	err = w.AddNamespace("api/v1/schedule").AddRoutes(
		// schedule
		web.NewRoute(web.MethodPost, "/api/v1/schedule", c.CreateSchedule),
		web.NewRoute(web.MethodGet, "/api/v1/schedule", c.ListSchedule),
		web.NewRoute(web.MethodGet, "/api/v1/schedule/:id", c.GetSchedule),
		web.NewRoute(web.MethodPut, "/api/v1/schedule/:id", c.UpdateSchedule),
		web.NewRoute(web.MethodDelete, "/api/v1/schedule/:id", c.DeleteSchedule),
	)

	return err
}
