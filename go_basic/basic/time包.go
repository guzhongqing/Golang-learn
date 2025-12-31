package main

import (
	"fmt"
	"time"
)

func basicTime() {
	t0 := time.Now()
	fmt.Println(t0.Unix())
	time.Sleep(2 * time.Second)
	t1 := time.Now()
	fmt.Println(t1.Unix())
	// 时间差t1-t0的时间
	fmt.Println(t1.Sub(t0).Seconds())
	fmt.Println("t1在t0之后吗？", t1.After(t0))

	t2 := t1.Add(2 * time.Hour)
	fmt.Println(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second())

	fmt.Println(t2.YearDay(), t2.Weekday(), t2.Weekday().String(), t2.Format("2006-01-02 15:04:05"))
}

// 周期执行
func ticker() {
	const TIME_FMT = "2006-01-02 15:04:05"
	fmt.Println("当前时间:", time.Now().Format(TIME_FMT))
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		<-tk.C
		fmt.Println("当前时间:", time.Now().Format(TIME_FMT))
	}
}

// 定时执行
func timer() {
	const TIME_FMT = "2006-01-02 15:04:05"
	fmt.Println("当前时间:", time.Now().Format(TIME_FMT))
	tm := time.NewTimer(5 * time.Second)
	<-tm.C
	fmt.Println("当前时间:", time.Now().Format(TIME_FMT))
	<-time.After(2 * time.Second)
	fmt.Println("当前时间:", time.Now().Format(TIME_FMT))

}

func main31() {
	// basicTime()
	// ticker()
	timer()

}
