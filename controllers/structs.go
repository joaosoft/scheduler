package controllers

type GetTimezoneRequest struct {
	Id string `json:"id" validate:"not-empty"`
}

type GetScheduleRequest struct {
	Id string `json:"id" validate:"not-empty"`
}

type CreateScheduleRequest struct {
	Body CreateScheduleBodyRequest `json:"body" validate:"not-empty"`
}

type CreateScheduleBodyRequest struct {
	Subject     string `json:"subject" validate:"not-empty"`
	Description string `json:"description"`
	IdUser      int    `json:"id_user" validate:"not-empty"`
}

type UpdateScheduleRequest struct {
	Id   string                    `json:"id" validate:"not-empty"`
	Body UpdateScheduleBodyRequest `json:"body" validate:"not-empty"`
}

type UpdateScheduleBodyRequest struct {
	Subject     string `json:"subject" validate:"not-empty"`
	Description string `json:"description"`
	IdTimezone  int    `json:"id_timezone" validate:"not-empty"`
}

type DeleteScheduleRequest struct {
	Id string `json:"id" validate:"not-empty"`
}
