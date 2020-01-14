package samples

import "github.com/gonutz/w32"

//TestIcon func
func TestIcon() {
	hwnd := w32.FindWindow("", "Безымянный — Блокнот")
	hIcon := w32.ExtractIcon("E:\\GOPROJECTS\\win_hack\\assets\\app.ico", 0)
	w32.SendMessage(hwnd, w32.WM_SETICON, w32.ICON_SMALL, uintptr(hIcon))
	w32.SendMessage(hwnd, w32.WM_SETICON, w32.ICON_SMALL2, uintptr(hIcon))
}
