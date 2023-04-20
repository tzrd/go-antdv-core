package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tzrd/go-antdv-core/timex"
)

// An ElapsedTimer is a timer to track the elapsed time.
type ElapsedTimer struct {
	start time.Duration
}

// NewElapsedTimer returns an ElapsedTimer.
func NewElapsedTimer() *ElapsedTimer {
	return &ElapsedTimer{
		start: timex.Now(),
	}
}

// Duration returns the elapsed time.
func (et *ElapsedTimer) Duration() time.Duration {
	return timex.Since(et.start)
}

// Elapsed returns the string representation of elapsed time.
func (et *ElapsedTimer) Elapsed() string {
	return timex.Since(et.start).String()
}

// ElapsedMs returns the elapsed time of string on milliseconds.
func (et *ElapsedTimer) ElapsedMs() string {
	return fmt.Sprintf("%.1fms", float32(timex.Since(et.start))/float32(time.Millisecond))
}

// CurrentMicros returns the current microseconds.
func CurrentMicros() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// CurrentMillis returns the current milliseconds.
func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Difference in days between two dates
func DateSubDays(format string, last, first string) (int, error) {
	date1, err := time.ParseInLocation(format, last, time.Local)
	if err != nil {
		return 0, err
	}

	date2, err := time.ParseInLocation(format, first, time.Local)
	if err != nil {
		return 0, err
	}

	return int(date1.Sub(date2).Hours() / 24), nil
}

// Formulate the number of days in the year and month
func YearMonthDays(year int, month int) int {
	//有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] {
		return 31
	}
	// 有30天的月份

	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] {
		return 30
	}
	//计算平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出二月天数
		return 29
	}
	// 得出平年二月天数
	return 28
}

func ElapsedDays(month int) (int, error) {
	var elapsed = 0
	var err error = nil

	ym := strconv.Itoa(month)
	now := time.Now()

	if ym < now.Format("200601") {
		year_str := fmt.Sprintf("%v", month)[:4]
		month_str := fmt.Sprintf("%v", month)[4:]
		y, e := strconv.Atoi(year_str)
		if e != nil {
			return -1, e
		}
		m, e := strconv.Atoi(month_str)
		if e != nil {
			return -1, e
		}
		elapsed = YearMonthDays(y, m)
	} else if ym == now.Format("200601") {
		first := time.Now().Format("200601") + "01"
		last := time.Now().Format("20060102")
		elapsed, err = DateSubDays("20060102", last, first)
		if err != nil {
			return -1, err
		}
	} else {
		elapsed = 0
	}

	return elapsed, err
}
