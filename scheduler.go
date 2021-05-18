package scheduler

import (
	"scheduler/routes"

	"github.com/joaosoft/logger"
	migration "github.com/joaosoft/migration/services"
	"github.com/joaosoft/web"
)

type Scheduler struct {
	web    *web.Server
	config *SchedulerConfig
	logger logger.ILogger
}

// NewScheduler ...
func NewScheduler(options ...SchedulerOption) (*Scheduler, error) {
	config, _, err := newConfig()

	service := &Scheduler{
		logger: logger.NewLogDefault("scheduler", logger.WarnLevel),
		config: config.Scheduler,
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Scheduler != nil {
		level, _ := logger.ParseLevel(config.Scheduler.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.Scheduler = &SchedulerConfig{
			Server: &web.ServerConfig{
				Address: defaultURL,
			},
		}
	}

	service.Reconfigure(options...)

	// execute migrations
	migrationService, err := migration.NewCmdService(migration.WithCmdConfiguration(service.config.Migration))
	if err != nil {
		return nil, err
	}

	if _, err = migrationService.Execute(migration.OptionUp, 0, migration.ExecutorModeDatabase); err != nil {
		return nil, err
	}

	service.web, err = web.NewServer(web.WithServerConfiguration(service.config.Server))
	if err != nil {
		return nil, err
	}

	ns := service.web.AddNamespace("/api/v1")

	// user
	if err = routes.RegisterUserRoutes(ns, service.logger, service.config.Dbr); err != nil {
		return nil, err
	}

	// timezone
	if err = routes.RegisterTimezoneRoutes(ns, service.logger, service.config.Dbr); err != nil {
		return nil, err
	}

	// schedule
	if err = routes.RegisterScheduleRoutes(ns, service.logger, service.config.Dbr); err != nil {
		return nil, err
	}

	return service, nil
}

// Start ...
func (s *Scheduler) Start() error {
	return s.web.Start()
}
