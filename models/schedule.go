package models

import (
	"time"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/logger"
)

type ScheduleModel struct {
	logger logger.ILogger
	db     *dbr.Dbr
}

func NewScheduleModel(logger logger.ILogger, config *dbr.DbrConfig) (*ScheduleModel, error) {
	db, err := dbr.New(dbr.WithConfiguration(config))
	if err != nil {
		return nil, err
	}

	return &ScheduleModel{
		logger: logger,
		db:     db,
	}, nil
}

func (m *ScheduleModel) ListSchedule() (scheduleList ScheduleList, err error) {
	_, err = m.db.
		Select([]interface{}{
			"s.id_schedule",
			"s.hashed_id",
			"s.subject",
			"s.description",
			"s.fk_timezone",
		}...).
		From(dbr.As(schedulerTableSchedule, "s")).
		Where("s.active").
		Load(&scheduleList)

	return scheduleList, err
}

func (m *ScheduleModel) GetSchedule(param *GetSchedule) (schedule *Schedule, err error) {
	var count int
	count, err = m.db.
		Select([]interface{}{
			"s.id_schedule",
			"s.hashed_id",
			"s.subject",
			"s.description",
			"s.fk_timezone",
		}...).
		From(dbr.As(schedulerTableSchedule, "s")).
		Where("s.id_schedule = ?", param.Id).
		Where("s.active").
		Load(schedule)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return schedule, nil
}

func (m *ScheduleModel) CreateSchedule(param *CreateSchedule) (schedule *Schedule, err error) {
	var tx *dbr.Transaction
	if tx, err = m.db.Begin(); err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommit()

	now := time.Now()
	schedule = &Schedule{}
	_, err = tx.
		Insert().
		Into(schedulerTableSchedule).
		Record(param).
		Return([]interface{}{
			"id_schedule",
			"hashed_id",
			"subject",
			"description",
			"fk_user",
			"fk_schedule_status",
		}...).
		Load(schedule)

	if err != nil {
		return nil, err
	}

	builder := tx.
		Insert().
		Columns([]interface{}{
			"fk_schedule",
			"time",
			"position",
			"active",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
		}...).
		Into(dbr.As(schedulerTableScheduleTimeSlot, "ts")).
		Return([]interface{}{
			"ts.fk_schedule",
			"ts.time",
		}...)

	for idx, timeSlot := range param.TimeSlots {
		schedule = &Schedule{}
		builder.Values(
			schedule.Id,
			timeSlot,
			idx+1,
			param.IdUser,
			now,
			param.IdUser,
			now,
		)
	}

	if _, err = builder.Load(schedule); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return schedule, nil
}

func (m *ScheduleModel) UpdateSchedule(param *UpdateSchedule) (schedule *Schedule, err error) {
	var count int
	count, err = m.db.
		Update(dbr.As(schedulerTableSchedule, "s")).
		Record(param).
		Return([]interface{}{
			"s.id_schedule",
			"s.hashed_id",
			"s.subject",
			"s.description",
			"s.fk_timezone",
		}...).
		Load(&schedule)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return schedule, nil
}

func (m *ScheduleModel) DeleteSchedule(param *DeleteSchedule) (err error) {
	var count int
	_, err = m.db.
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
