package date_utils

import (
	"time"
)

const (
	defaultDateLayout = "2006-01-02T15:04:05Z"
	dbDateLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(defaultDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(dbDateLayout)
}

