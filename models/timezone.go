package models

import (
	"github.com/joaosoft/dbr"
	"github.com/joaosoft/logger"
)

type TimezoneModel struct {
	logger logger.ILogger
	db     *dbr.Dbr
}

func NewTimezoneModel(logger logger.ILogger, config *dbr.DbrConfig) (*TimezoneModel, error) {
	db, err := dbr.New(dbr.WithConfiguration(config))
	if err != nil {
		return nil, err
	}

	return &TimezoneModel{
		logger: logger,
		db:     db,
	}, nil
}

func (m *TimezoneModel) ListTimezone() (timezoneList TimezoneList, err error) {
	_, err = m.db.
		Select([]interface{}{
			"t.id_timezone",
			"t.name",
		}...).
		From(dbr.As(schedulerTableTimezone, "t")).
		Where("t.active").
		Load(&timezoneList)

	if err != nil {
		return nil, err
	}

	return timezoneList, err
}

func (m *TimezoneModel) GetTimezone(param *GetTimezone) (timezone *Timezone, err error) {
	var count int
	count, err = m.db.
		Select([]interface{}{
			"t.id_timezone",
			"t.name",
		}...).
		From(dbr.As(schedulerTableTimezone, "t")).
		Where("t.id_timezone = ?", param.Id).
		Where("t.active").
		Load(timezone)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return timezone, err
}
