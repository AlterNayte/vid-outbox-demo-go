package internal

import "time"

func GetStringValue(s *string) *string {
	if s == nil {
		return nil
	}
	value := *s
	return &value
}

func GetTimeValue(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}
	value := *t
	return &value
}
