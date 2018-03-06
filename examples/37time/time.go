package main

import (
	"fmt"
	"time"
)

// 演示time
var p = fmt.Println

func main() {

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
