package main

import "strconv"

//var window = oc.window(30, 4, 40, 17)
var windows []windowStruct

type windowStruct struct {
	object
	visible        bool
	windowType     string
	additionalInfo string
}

func (w *windowStruct) init() {

	for i := 0; i < w.width*w.height; i++ {
		w.data[i] = ' '
	}

	for i := 1; i < w.height; i++ {
		w.data[i*w.width] = '|'
		w.data[i*w.width-1] = '|'
	}
	for i := 0; i < w.width; i++ {
		w.data[i] = '-'
		if i != 0 && i != w.width-1 {
			w.data[w.width+i] = '-'
		}
		w.data[len(w.data)-1-i] = '-'
	}
	w.data[2*w.width-2] = 'X'

	if w.windowType == "help" {
		w.writeTxt(1, 2, "Help")
		w.writeTxt(1, 4, "Use o to upen up a window with the")
		w.writeTxt(2, 5, "parameters n for notepad and h for")
		w.writeTxt(2, 6, "help, use c to close it.")
		w.writeTxt(1, 7, "Switch between the windows using s")
		w.writeTxt(2, 8, "with the optional number of the")
		w.writeTxt(2, 9, "window to switch to.")
		w.writeTxt(1, 10, "Move the window with m in the")
		w.writeTxt(2, 11, "directions u d l and r.")
		w.writeTxt(1, 12, "Use d to show the desktop.")
		w.writeTxt(1, 13, "Use w to write into the notepad.")
	}
	if w.windowType == "notepad" {
		w.additionalInfo = "2"
	}
}

func (w *windowStruct) sendInfo(info string) {
	if w.windowType == "notepad" {
		if len(info) > w.width-2 {
			info = info[0 : w.width-2]
		}
		lineNumber, _ := strconv.Atoi(w.additionalInfo)
		w.writeTxt(1, lineNumber, info)
		lineNumber++
		w.additionalInfo = strconv.Itoa(lineNumber)
	}
}
