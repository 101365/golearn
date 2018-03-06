package main

import (
	"fmt"
	"time"
)

// 演示time 时间 以及时间戳
var p = fmt.Println

func main() {

	// 演示时间
	showTime()

	// 演示时间戳
	showTimestamp()

	// 演示时间格式化
	formatTimeDemo()

}

// 演示时间格式化
func formatTimeDemo() {

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

}

// 演示时间戳
func showTimestamp() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

}

// 演示时间
func showTime() {
	// 当前时间
	now := time.Now()
	p(now)

	// 指定时间
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	// 和当前时间比较
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 相差时间范围
	diff := now.Sub(then)
	p("diff => ", diff)

	// 相差的小时、分钟、秒、毫秒
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 时间加减
	p(then.Add(diff))
	p(then.Add(-diff))
}
