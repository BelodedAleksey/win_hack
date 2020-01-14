package samples

import (
	"github.com/gonutz/w32"
	"main/win"
	"syscall"

	"unsafe"
)

//TestText func
func TestText() {
	hwnd := w32.FindWindow("", "Безымянный — Блокнот")
	hEdit := win.FindWindowEx(hwnd, 0, "Edit", "")
	w32.SendMessage(
		hEdit, w32.EM_SETSEL, 0, 0,
	)
	w32.SendMessage(
		hEdit, w32.EM_REPLACESEL, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("TEXT"))),
	)
}
