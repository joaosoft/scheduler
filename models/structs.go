package models

type GetTimezone struct {
	Id string `json:"id"`
}

type GetSchedule struct {
	Id string `json:"id"`
}

type CreateSchedule struct {
	HashedId         string `json:"hashed_id" db.write:"hashed_id"`
	Subject          string `json:"subject" db.write:"subject"`
	Description      string `json:"description" db.write:"description"`
	IdUser           int    `json:"id_user" db.write:"fk_user"`
	IdScheduleStatus int    `json:"id_schedule_status" db.write:"fk_schedule_status"`
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
	IdUser      int    `json:"id_user" db.read:"fk_user"`
}
