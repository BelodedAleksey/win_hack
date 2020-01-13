package samples

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"main/win"

	"github.com/go-vgo/robotgo"
	"github.com/gonutz/w32"

	"github.com/sirupsen/logrus"
)

const (
	targetX = 100
	targetY = 100
)

var moved bool
var oldX, oldY int32

//LowLevelMouseProcess ...
func LowLevelMouseProcess(nCode int, wparam w32.WPARAM, lparam w32.LPARAM) w32.LRESULT {
	var temporaryKeyPtr w32.HHOOK
	var mouse *win.MSLLHOOKSTRUCT
	if nCode == 0 && wparam == w32.WM_MOUSEMOVE {
		mouse = (*win.MSLLHOOKSTRUCT)(unsafe.Pointer(lparam))
		x := mouse.Pt.X
		y := mouse.Pt.Y
		logrus.Infoln("X: ", x, "Y: ", y)
		fmt.Println("X: ", x, "Y: ", y)
	}
	return w32.CallNextHookEx(temporaryKeyPtr, nCode, wparam, lparam)
}

//TestMouse func
func TestMouse() {
	// Set Log output to a file
	filename := fmt.Sprintf("C:\\mouse.txt")
	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Info("Can't set log file")
	}
	logrus.SetOutput(logFile)
	//Start()
	var msg w32.MSG
	mouseHook := w32.SetWindowsHookEx(w32.WH_MOUSE_LL, LowLevelMouseProcess, 0, 0)
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
	}
	w32.UnhookWindowsHookEx(mouseHook)
}

//TestMouseCircle func
func TestMouseCircle() {
	//mouseStart()
	var oldX, oldY int
	var targetX, targetY int
	const r = 300
	for {
		x, y := robotgo.GetMousePos()
		if math.Abs(float64(oldX-x)) > 5 || math.Abs(float64(oldY-y)) > 5 {
			oldX = x
			oldY = y
			//Circle
			for i := 0.0; i < 2.0; i = i + 0.15 {
				targetX = int(math.Sin(i*math.Pi)*r) + 800
				targetY = int(math.Cos(i*math.Pi)*r) + 500
				robotgo.MoveMouseSmooth(targetX, targetY)
			}
		}
	}
}
