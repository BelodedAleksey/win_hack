package win

import (
	"syscall"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	gdi32    = syscall.NewLazyDLL("gdi32.dll")

	LoadLibraryProc = kernel32.NewProc("LoadLibraryW")

	createPatternBrushProc = gdi32.NewProc("CreatePatternBrush")

	messageBoxProc                 = user32.NewProc("MessageBoxW")
	setLayeredWindowAttributesProc = user32.NewProc("SetLayeredWindowAttributes")
	procToUnicode                  = user32.NewProc("ToUnicode")
	procGetKeyboardState           = user32.NewProc("GetKeyboardState")
	procGetKeyboardLayoutName      = user32.NewProc("GetKeyboardLayoutNameW")
	procGetKeyboardLayout          = user32.NewProc("GetKeyboardLayout")
	enumWindowsProc                = user32.NewProc("EnumWindows")
	setActiveWindowProc            = user32.NewProc("SetActiveWindow")
	findWindowExProc               = user32.NewProc("FindWindowExW")
)
