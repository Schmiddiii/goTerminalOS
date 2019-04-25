package main

type desktopStruct struct {
	object
	taskbar taskbarStruct
}

var desktop = oc.desktop()
