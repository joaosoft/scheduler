package scheduler

import (
	"fmt"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
	"github.com/joaosoft/web"
)

// AppConfig ...
type AppConfig struct {
	Scheduler *SchedulerConfig `json:"scheduler"`
}

// SchedulerConfig ...
type SchedulerConfig struct {
	Server    *web.ServerConfig          `json:"server"`
	Dbr       *dbr.DbrConfig             `json:"dbr"`
	Migration *migration.MigrationConfig `json:"migration"`
	Log       struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func newConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
