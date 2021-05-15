package models

type GetTimezone struct {
	Id string `json:"id"`
}

type GetSchedule struct {
	Id string `json:"id"`
}

type CreateSchedule struct {
	Subject     string `json:"subject" db.write:"subject"`
	Description string `json:"description" db.write:"description"`
	IdTimezone  int    `json:"id_timezone" db.write:"fk_timezone"`
}

type UpdateSchedule struct {
	Subject     string `json:"subject" db.write:"subject"`
	Description string `json:"description" db.write:"description"`
	IdTimezone  int    `json:"id_timezone" db.write:"fk_timezone"`
}

type DeleteSchedule struct {
	Id string `json:"id"`
}

type TimezoneList []*Timezone

type Timezone struct {
	Id   string `json:"id" db.read:"id_timezone"`
	Name string `json:"name" db.read:"name"`
}

type ScheduleList []*Schedule

type Schedule struct {
	Id          string `json:"id" db.read:"id_schedule"`
	Subject     string `json:"subject" db.read:"subject"`
	Description string `json:"description" db.read:"description"`
	IdTimezone  int    `json:"id_timezone" db.read:"fk_timezone"`
}
