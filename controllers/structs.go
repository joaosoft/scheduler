package controllers

import "time"

type GetTimezoneRequest struct {
	Id string `json:"id" validate:"not-empty"`
}

type GetUserRequest struct {
	Id string `json:"id" validate:"not-empty"`
}

type CreateUserRequest struct {
	Body CreateUserBodyRequest `json:"body" validate:"not-empty"`
}

type CreateUserBodyRequest struct {
	Id         int `json:"id" validate:"not-empty"`
	IdCountry  int `json:"id_country" validate:"not-empty"`
	IdTimezone int `json:"id_timezone" validate:"not-empty"`
}

type UpdateUserRequest struct {
	Id   string                `json:"id" validate:"not-empty"`
	Body UpdateUserBodyRequest `json:"body" validate:"not-empty"`
}

type UpdateUserBodyRequest struct {
	Id         int `json:"id" validate:"not-empty"`
	IdCountry  int `json:"id_country" validate:"not-empty"`
	IdTimezone int `json:"id_timezone" validate:"not-empty"`
}

type DeleteUserRequest struct {
	Id string `json:"id" validate:"not-empty"`
}

type GetScheduleRequest struct {
	Id string `json:"id" validate:"not-empty"`
}

type CreateScheduleRequest struct {
	Body CreateScheduleBodyRequest `json:"body" validate:"not-empty"`
}

type CreateScheduleBodyRequest struct {
	IdUser      int         `json:"id_user" validate:"not-empty"`
	IdTimezone  int         `json:"id_timezone" validate:"not-empty"`
	Subject     string      `json:"subject" validate:"not-empty"`
	Description string      `json:"description"`
	TimeSlots   []time.Time `json:"time_slots"`
}

type UpdateScheduleRequest struct {
	Id   string                    `json:"id" validate:"not-empty"`
	Body UpdateScheduleBodyRequest `json:"body" validate:"not-empty"`
}

type UpdateScheduleBodyRequest struct {
	IdUser      int    `json:"id_user" validate:"not-empty"`
	IdTimezone  int    `json:"id_timezone" validate:"not-empty"`
	Subject     string `json:"subject" validate:"not-empty"`
	Description string `json:"description"`
}

type DeleteScheduleRequest struct {
	Id string `json:"id" validate:"not-empty"`
}
