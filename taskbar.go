package main

import (
	"strconv"
	"time"
)

type taskbarStruct struct {
	object
}

func (t *taskbarStruct) init() {
	for i := 0; i < t.width*t.height; i++ {
		t.data[i] = ' '
	}

	for i := 0; i < width; i++ {
		t.set(i, 0, '-')
	}

	//Symbol
	t.writeTxt(1, 1, "##")
	t.writeTxt(1, 2, "##")
	t.set(4, 1, '|')
	t.set(4, 2, '|')

	//Time
	t.set(width-13, 1, '|')
	t.set(width-13, 2, '|')
}

func (t *taskbarStruct) drawTimeDate() {
	t.drawDate()
	t.drawTime()
}

func (t *taskbarStruct) drawTime() {
	now := time.Now()
	formated := now.Format("15:04")

	t.writeTxt(width-8, 1, formated)

}

func (t *taskbarStruct) drawDate() {
	now := time.Now()

	year, month, day := now.Date()
	strDay := strconv.Itoa(day)
	if len(strDay) == 1 {
		strDay = "0" + strDay
	}
	strMonth := strconv.Itoa(int(month))
	if len(strMonth) == 1 {
		strMonth = "0" + strMonth
	}
	strYear := strconv.Itoa(year)

	t.writeTxt(width-11, 2, strDay+"."+strMonth+"."+strYear)

}
