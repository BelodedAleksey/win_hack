package win

import (
	"syscall"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	LoadLibraryProc                = kernel32.NewProc("LoadLibraryW")
	messageBoxProc                 = user32.NewProc("MessageBoxW")
	setLayeredWindowAttributesProc = user32.NewProc("SetLayeredWindowAttributes")
	procToUnicode                  = user32.NewProc("ToUnicode")
	procGetKeyboardState           = user32.NewProc("GetKeyboardState")
	procGetKeyboardLayoutName      = user32.NewProc("GetKeyboardLayoutNameW")
	procGetKeyboardLayout          = user32.NewProc("GetKeyboardLayout")
)
