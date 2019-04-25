package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxRefreshRate = 500 * time.Millisecond
const width, height = 100, 27

var oc = objectCreator{}

func main() {

	windows = append(windows, oc.help(30, 4, 40, 17))

	for {
		desktop.taskbar.drawTimeDate()
		draw()
		in := getInput()

		if in == "shutdown" {
			drawShutdown()
		}

		switch in[0] {
		case 'o':
			switch in[2] {
			case 'h':
				windows = append(windows, oc.help(len(windows), len(windows), 40, 17))
			case 'n':
				windows = append(windows, oc.notepad(len(windows), len(windows), 40, 17))
			}
			break
		case 's':
			if len(in) < 3 {
				if len(windows) > 1 {
					storeW := windows[len(windows)-1]
					windows[len(windows)-1] = windows[len(windows)-2]
					windows[len(windows)-2] = storeW
				}
				windows[len(windows)-1].visible = true
				break
			}
			index, _ := strconv.Atoi(string(in[2]))
			if index > 0 && index <= len(windows) {
				storeW := windows[len(windows)-1]
				windows[len(windows)-1] = windows[len(windows)-index]
				windows[len(windows)-index] = storeW
			}

		case 'd':
			for i := 0; i < len(windows); i++ {
				windows[i].visible = false
			}
			break
		case 'w':
			windows[len(windows)-1].sendInfo(string(in[2:len(in)]))
		case 'm':
			switch in[2] {
			case 'u':
				windows[len(windows)-1].y = 0
				break
			case 'd':
				windows[len(windows)-1].y = height - windows[len(windows)-1].height - desktop.taskbar.height
				break
			case 'l':
				windows[len(windows)-1].x = 0
				break
			case 'r':
				windows[len(windows)-1].x = width - windows[len(windows)-1].width
				break
			}
		case 'c':
			windows = windows[0 : len(windows)-1]
			break
		}

	}
}

func drawShutdown() {
	draw := make([]rune, width*height)

	draw[getIndex(width/2, height/2)-4] = 'G'
	draw[getIndex(width/2, height/2)-3] = 'O'
	draw[getIndex(width/2, height/2)-2] = 'O'
	draw[getIndex(width/2, height/2)-1] = 'D'
	draw[getIndex(width/2, height/2)] = 'B'
	draw[getIndex(width/2, height/2)+1] = 'Y'
	draw[getIndex(width/2, height/2)+2] = 'E'

	//Top side of monitor
	for x := 0; x < width+2; x++ {
		fmt.Print("-")
	}

	//Actually drawing out
	for i := 0; i < len(draw); i++ {
		//Lineending
		if i%width == 0 {
			if i != 0 {
				fmt.Print("|")
			}
			fmt.Println()
			fmt.Print("|")
		}

		fmt.Print(string(draw[i]))
	}
	//Final line ending is needed
	fmt.Print("|")
	fmt.Println()

	for x := 0; x < width+2; x++ {
		fmt.Print("-")
	}
	fmt.Println()

	time.Sleep(2 * time.Second)

	os.Exit(0)
}

func draw() {
	draw := make([]rune, width*height)

	//adding desktop to be drawn out
	draw = desktop.drawOnTopOf(draw)
	draw = desktop.taskbar.drawOnTopOf(draw)

	for _, w := range windows {
		if w.visible {
			draw = w.drawOnTopOf(draw)
		}

	}

	//Top side of monitor
	for x := 0; x < width+2; x++ {
		fmt.Print("-")
	}

	//Actually drawing out
	for i := 0; i < len(draw); i++ {
		//Lineending
		if i%width == 0 {
			if i != 0 {
				fmt.Print("|")
			}
			fmt.Println()
			fmt.Print("|")
		}

		fmt.Print(string(draw[i]))
	}
	//Final line ending is needed
	fmt.Print("|")
	fmt.Println()

	for x := 0; x < width+2; x++ {
		fmt.Print("-")
	}
	fmt.Println()

	time.Sleep(maxRefreshRate)
}

func getInput() string {
	var reader = bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	//Deleting the end of the line
	text = strings.Replace(text, "\r\n", "", -1)
	return text
}

func getIndex(x, y int) int {
	return y*width + x
}

func getIndexWidth(x, y, w int) int {
	return y*w + x
}
