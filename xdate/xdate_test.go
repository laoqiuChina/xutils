package xdate

import (
	"github.com/gogf/gf/g/os/gtime"
	"testing"
	"time"
)

func TestNowUnix(t *testing.T) {
	t.Log("TestNowUnix Begin")
	t.Log(NowUnix())
	t.Log("TestNowUnix End")
}

func TestNowTimestamp(t *testing.T) {
	t.Log("TestNowTimestamp Begin")
	t.Log(NowTimestamp())
	t.Log("TestNowTimestamp End")
}

func TestTimeUnix(t *testing.T) {
	t.Log("TestTimeUnix Begin")
	time := time.Now()
	long := TimeUnix(time)
	oldTime := TimeFromUnix(long)
	t.Log(time)
	t.Log(long)
	t.Log(oldTime)
	t.Log("TestTimeUnix End")
}

func TestTimestamp(t *testing.T) {
	t.Log("TestTimestamp Begin")
	time := time.Now()
	long := Timestamp(time)
	oldTime := TimeFromTimestamp(long)
	t.Log(time)
	t.Log(long)
	t.Log(oldTime)
	t.Log("TestTimestamp End")
}
func TestGetDay(t *testing.T) {
	t.Log("TestGetDay Begin")
	t.Log(GetDay(time.Now()))
	t.Log(GetDayLayout(time.Now(), FMT_DATE_TIME))
	t.Log("TestGetDay End")
}
func TestTimeParse(t *testing.T) {
	thisDate := "2014-03-17 14:55:06"
	t.Log(TimeParse(thisDate, FMT_DATE_TIME))
	t.Log(TimeParse(thisDate, FMT_DATE))
}
func TestWithTimeAsOfDay(t *testing.T) {
	t.Log("TestWithTimeAsStartOfDay Begin")
	t.Log(WithTimeAsStartOfDay(time.Now()))
	t.Log(WithTimeAsEndOfDay(time.Now()))
	t.Log("TestWithTimeAsStartOfDay End")
}

func TestXTime(t *testing.T) {
	x := xtime{}
	x.time = time.Now().Add(-1000 * time.Second)
	t.Log(x.time)
	t.Log(x.WithTimeAsStartOfDay())
	t.Log(x.WithTimeAsEndOfDay())
	t.Log(x.PrettyTime())
	x.time = time.Now().Add(-1000 * time.Minute)
	t.Log(x.PrettyTime())
	x.time = time.Now().Add(-25 * time.Hour)
	t.Log(x.PrettyTime())
	x.time = time.Now().Add(-47 * time.Hour)
	t.Log(x.PrettyTime())
	x.time = time.Now().Add(-73 * time.Hour)
	t.Log(x.PrettyTime())
	t.Log(x.TimeFormat(FMT_DATE_TIME_CN))
}

func TestTimeFormat(t *testing.T) {
	t.Log(TimeFormat(gtime.Now().Time, FMT_DATE_TIME_CN))
}
