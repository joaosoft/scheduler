package models

import (
	"github.com/joaosoft/dbr"
)

type StoragePostgres struct {
	db *dbr.Dbr
}

func NewStoragePostgres(config *dbr.DbrConfig) (*StoragePostgres, error) {
	dbr, err := dbr.New(dbr.WithConfiguration(config))
	if err != nil {
		return nil, err
	}

	return &StoragePostgres{
		db: dbr,
	}, nil
}

func (s *StoragePostgres) ListTimezone() (timezoneList TimezoneList, err error) {
	_, err = s.db.
		Select([]interface{}{"t.id_timezone", "t.name"}...).
		From(dbr.As(schedulerTableTimezone, "t")).
		Where("t.active").
		Load(&timezoneList)

	if err != nil {
		return nil, err
	}

	return timezoneList, nil
}

func (s *StoragePostgres) GetTimezone(param *GetTimezone) (timezone *Timezone, err error) {
	var count int
	count, err = s.db.
		Select([]interface{}{"t.id_timezone", "t.name"}...).
		From(dbr.As(schedulerTableTimezone, "t")).
		Where("t.id_timezone = ?", param.Id).
		Where("t.active").
		Load(&timezone)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return timezone, nil
}

func (s *StoragePostgres) ListSchedule() (scheduleList ScheduleList, err error) {
	_, err = s.db.
		Select([]interface{}{"s.id_schedule", "s.subject", "s.description", "s.fk_timezone"}...).
		From(dbr.As(schedulerTableSchedule, "s")).
		Where("s.active").
		Load(&scheduleList)

	if err != nil {
		return nil, err
	}

	return scheduleList, nil
}

func (s *StoragePostgres) GetSchedule(param *GetSchedule) (schedule *Schedule, err error) {
	var count int
	count, err = s.db.
		Select([]interface{}{"s.id_schedule", "s.subject", "s.description", "s.fk_timezone"}...).
		From(dbr.As(schedulerTableSchedule, "s")).
		Where("s.id_schedule = ?", param.Id).
		Where("s.active").
		Load(&schedule)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return schedule, nil
}

func (s *StoragePostgres) CreateSchedule(param *CreateSchedule) (schedule *Schedule, err error) {
	_, err = s.db.
		Insert().
		Into(dbr.As(schedulerTableSchedule, "s")).
		Record(param).
		Return([]interface{}{"s.id_schedule", "s.subject", "s.description", "s.fk_timezone"}...).
		Load(&schedule)

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (s *StoragePostgres) UpdateSchedule(param *UpdateSchedule) (schedule *Schedule, err error) {
	var count int
	count, err = s.db.
		Update(dbr.As(schedulerTableSchedule, "s")).
		Record(param).
		Return([]interface{}{"s.id_schedule", "s.subject", "s.description", "s.fk_timezone"}...).
		Load(&schedule)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return schedule, nil
}

func (s *StoragePostgres) DeleteSchedule(param *DeleteSchedule) (err error) {
	var count int
	_, err = s.db.
		Delete().
		From(dbr.As(schedulerTableSchedule, "s")).
		Where("id_schedule = ?", param.Id).
		Exec()

	if err != nil {
		return err
	}

	if count == 0 {
		return ErrorNotFound
	}

	return nil
}
