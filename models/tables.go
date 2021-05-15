package models

import (
	"fmt"
)

func format(schema, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

var (
	schedulerTableSchedule = format(schemaScheduler, "schedule")
	schedulerTableTimezone = format(schemaScheduler, "timezone")
)
