package models

import (
	"github.com/joaosoft/logger"
)

type IStorageDB interface {
	ListTimezone() (TimezoneList, error)
	GetTimezone(param *GetTimezone) (*Timezone, error)
	ListSchedule() (ScheduleList, error)
	GetSchedule(param *GetSchedule) (*Schedule, error)
	CreateSchedule(param *CreateSchedule) (*Schedule, error)
	UpdateSchedule(param *UpdateSchedule) (*Schedule, error)
	DeleteSchedule(param *DeleteSchedule) error
}

type Model struct {
	logger  logger.ILogger
	storage IStorageDB
}

func NewModel(logger logger.ILogger, storageDB IStorageDB) *Model {
	return &Model{
		logger:  logger,
		storage: storageDB,
	}
}

func (m *Model) ListTimezone() (timezoneList TimezoneList, err error) {
	timezoneList, err = m.storage.ListTimezone()
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	return timezoneList, err
}

func (m *Model) GetTimezone(param *GetTimezone) (timezone *Timezone, err error) {
	timezone, err = m.storage.GetTimezone(param)
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	return timezone, err
}

func (m *Model) ListSchedule() (scheduleList ScheduleList, err error) {
	scheduleList, err = m.storage.ListSchedule()
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	return scheduleList, err
}

func (m *Model) GetSchedule(param *GetSchedule) (schedule *Schedule, err error) {
	schedule, err = m.storage.GetSchedule(param)
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	return schedule, err
}

func (m *Model) CreateSchedule(param *CreateSchedule) (schedule *Schedule, err error) {
	schedule, err = m.storage.CreateSchedule(param)
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	return schedule, err
}

func (m *Model) UpdateSchedule(param *UpdateSchedule) (schedule *Schedule, err error) {
	schedule, err = m.storage.UpdateSchedule(param)
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	return schedule, err
}

func (m *Model) DeleteSchedule(param *DeleteSchedule) error {
	err := m.storage.DeleteSchedule(param)
	if err != nil {
		m.logger.Error(err)
		return err
	}

	return nil
}
