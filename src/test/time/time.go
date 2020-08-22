package main

import (
	"log"
	"time"
)

func main() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	inputTime := "2029-09-04 12:02:33"
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, inputTime)
	dateTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
	log.Printf("输入时间: %s, 输出时间: %s", inputTime, dateTime)

	/**
	*从输出结果来看，输入时间和输出时间竟然相差了八个小时，这显然是时区的设置问题。
	*在调用Format方法前我们已经设置了时区，为什么还会出现时区问题呢？
	*实际上这与Parse方法有直接关系，因为Parse方法会尝试在入参的参数中分析并读取时区信息。
	*如果入参的参数没有指定时区信息，那么会默认使用UTC时间。
	*因此在这种情况下，我们采用ParseInLocation方法指定时区，就可以解决这个问题
	**/
	t, _ = time.ParseInLocation(layout, inputTime, location)
	dateTime = time.Unix(t.Unix(), 0).In(location).Format(layout)
	log.Printf("输入时间: %s, 输出时间: %s", inputTime, dateTime)
}
