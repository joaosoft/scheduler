package models

import "time"

type GetTimezone struct {
	Id string `json:"id"`
}

type GetUser struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Id         int `json:"id" db.write:"fk_user"`
	IdCountry  int `json:"id_country" db.write:"fk_country"`
	IdTimezone int `json:"id_timezone" db.write:"fk_timezone"`
}

type UpdateUser struct {
	Id         int `json:"id" db.write:"fk_user"`
	IdCountry  int `json:"id_country" db.write:"fk_country"`
	IdTimezone int `json:"id_timezone" db.write:"fk_timezone"`
}

type DeleteUser struct {
	Id string `json:"id"`
}

type User struct {
	Id         string `json:"id" db.read:"id_timezone"`
	IdCountry  string `json:"id_country" db.read:"fk_country"`
	IdTimezone string `json:"id_timezone" db.read:"fk_timezone"`
}

type GetSchedule struct {
	Id string `json:"id"`
}

type CreateSchedule struct {
	HashedId         string      `json:"hashed_id" db.write:"hashed_id"`
	Subject          string      `json:"subject" db.write:"subject"`
	Description      string      `json:"description" db.write:"description"`
	IdUser           int         `json:"id_user" db.write:"fk_user"`
	IdTimezone       int         `json:"id_timezone" db.write:"fk_timezone"`
	TimeSlots        []time.Time `json:"time_slots"`
	IdScheduleStatus int         `json:"id_schedule_status" db.write:"fk_schedule_status"`
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
	HashedId    string `json:"hashed_id" db.read:"hashed_id"`
	Subject     string `json:"subject" db.read:"subject"`
	Description string `json:"description" db.read:"description"`
	IdUser      int    `json:"id_user" db.read:"fk_user"`
}
