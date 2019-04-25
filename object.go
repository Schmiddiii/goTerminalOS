package main

type object struct {
	data          []rune
	width, height int
	x, y          int
}

func (o *object) init() {
	for i := 0; i < o.width*o.height; i++ {
		o.data[i] = ' '
	}
}

func (o *object) set(x, y int, c rune) {
	o.data[getIndexWidth(x, y, o.width)] = c
}

func (o *object) get(x, y int) rune {
	return o.data[getIndexWidth(x, y, o.width)]
}

func (o *object) writeTxt(x, y int, txt string) {
	for index, char := range txt {
		o.set(x+index, y, char)
	}
}

func (o *object) drawOnTopOf(str []rune) []rune {
	for iy := 0; iy < o.height; iy++ {
		for ix := 0; ix < o.width; ix++ {
			str[getIndex(o.x+ix, o.y+iy)] = o.get(ix, iy)
		}
	}
	return str
}

type objectCreator struct {
}

func (oc *objectCreator) desktop() desktopStruct {
	w := desktopStruct{object{data: make([]rune, width*height), width: width, height: height, x: 0, y: 0}, oc.taskbar()}
	w.init()

	return w
}

func (oc *objectCreator) taskbar() taskbarStruct {
	tb := taskbarStruct{object{data: make([]rune, 3*width), width: width, height: 3, x: 0, y: height - 3}}
	tb.init()

	return tb
}

func (oc *objectCreator) window(xW, yW, widthW, heightW int) windowStruct {
	w := windowStruct{object{data: make([]rune, widthW*heightW), width: widthW, height: heightW, x: xW, y: yW}, true, "window", ""}
	w.init()

	return w
}

func (oc *objectCreator) notepad(xW, yW, widthW, heightW int) windowStruct {
	w := windowStruct{object{data: make([]rune, widthW*heightW), width: widthW, height: heightW, x: xW, y: yW}, true, "notepad", ""}
	w.init()

	return w
}

func (oc *objectCreator) help(xW, yW, widthW, heightW int) windowStruct {
	w := windowStruct{object{data: make([]rune, widthW*heightW), width: widthW, height: heightW, x: xW, y: yW}, true, "help", ""}
	w.init()

	return w
}
