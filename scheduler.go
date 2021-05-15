package scheduler

import (
	"github.com/joaosoft/logger"
	migration "github.com/joaosoft/migration/services"
	"github.com/joaosoft/web"
	"scheduler/controllers"
	"scheduler/models"
	"scheduler/routes"
)

type Scheduler struct {
	web *web.Server
	config        *SchedulerConfig
	logger        logger.ILogger
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
			Host: defaultURL,
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

	storage, err := models.NewStoragePostgres(config.Scheduler.Dbr)
	if err != nil {
		return nil, err
	}

	model := models.NewModel(service.logger, storage)

	controller := controllers.NewController(model)
	service.web, err = web.NewServer()
	if err != nil {
		return nil, err
	}

	if err = routes.RegisterRoutes(service.web, controller); err != nil {
		return nil, err
	}

	return service, nil
}

// Start ...
func (s *Scheduler) Start() error {
	return s.web.Start()
}