package models

import (
	"fmt"
)

func format(schema, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

var (
	schedulerTableUser     = format(schemaScheduler, "user")
	schedulerTableTimezone = format(schemaScheduler, "timezone")
	schedulerTableSchedule = format(schemaScheduler, "schedule")
)
