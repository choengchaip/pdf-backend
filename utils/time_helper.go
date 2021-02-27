package utils

import "time"

func NewTimeStamp(t time.Time) time.Time {
	return t
}

func NewTimeStampS(t time.Time) string {
	return t.Format(time.RFC3339)
}

func NewTimeStampT(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, str)
	return t
}
