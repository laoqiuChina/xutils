package xdate

import "time"

type xtime struct {
	time time.Time
}

func (t xtime) WithTimeAsStartOfDay() time.Time {
	return WithTimeAsStartOfDay(t.time)
}
func (t xtime) WithTimeAsEndOfDay() time.Time {
	return WithTimeAsEndOfDay(t.time)
}

func (t xtime) PrettyTime() string {
	timestamp := Timestamp(t.time)
	return PrettyTime(timestamp)
}

func (t xtime) TimeFormat(layout string) string {
	return TimeFormat(t.time, layout)
}
