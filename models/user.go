package models

import (
	"github.com/joaosoft/dbr"
	"github.com/joaosoft/logger"
)

type UserModel struct {
	logger logger.ILogger
	db     *dbr.Dbr
}

func NewUserModel(logger logger.ILogger, config *dbr.DbrConfig) (*UserModel, error) {
	db, err := dbr.New(dbr.WithConfiguration(config))
	if err != nil {
		return nil, err
	}

	return &UserModel{
		logger: logger,
		db:     db,
	}, nil
}

func (m *UserModel) GetUser(param *GetUser) (user *User, err error) {
	user = &User{}
	var count int
	count, err = m.db.
		Select([]interface{}{
			"fk_user",
			"fk_country",
			"fk_timezone",
		}...).
		From(schedulerTableUser).
		Where("fk_user = ?", param.Id).
		Where("active").
		Load(user)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return user, nil
}

func (m *UserModel) CreateUser(param *CreateUser) (user *User, err error) {
	user = &User{}
	_, err = m.db.
		Insert().
		Into(schedulerTableUser).
		Record(param).
		Return([]interface{}{
			"fk_user",
			"fk_country",
			"fk_timezone",
		}...).
		Load(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UserModel) UpdateUser(param *UpdateUser) (user *User, err error) {
	var count int
	count, err = m.db.
		Update(schedulerTableUser).
		Record(param).
		Where("fk_user = ?", param.Id).
		Return([]interface{}{
			"fk_user",
			"fk_country",
			"fk_timezone",
		}...).
		Load(&user)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return user, nil
}

func (m *UserModel) DeleteUser(param *DeleteUser) (err error) {
	_, err = m.db.
		Delete().
		From(schedulerTableUser).
		Where("fk_user = ?", param.Id).
		Exec()

	if err != nil {
		return err
	}

	return nil
}
