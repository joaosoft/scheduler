package controllers

type GetTimezoneRequest struct {
	Id string `json:"id" validate:"notzero"`
}

type GetScheduleRequest struct {
	Id string `json:"id" validate:"notzero"`
}

type CreateScheduleRequest struct {
	Body CreateScheduleBodyRequest `json:"body" validate:"notzero"`
}

type CreateScheduleBodyRequest struct {
	Subject     string `json:"subject" validate:"notzero"`
	Description string `json:"description"`
	IdTimezone  int    `json:"id_timezone" validate:"notzero"`
}

type UpdateScheduleRequest struct {
	Id   string            `json:"id" validate:"notzero"`
	Body UpdateScheduleBodyRequest `json:"body" validate:"notzero"`
}

type UpdateScheduleBodyRequest struct {
	Subject     string `json:"subject" validate:"notzero"`
	Description string `json:"description"`
	IdTimezone  int    `json:"id_timezone" validate:"notzero"`
}

type DeleteScheduleRequest struct {
	Id string `json:"id" validate:"notzero"`
}
