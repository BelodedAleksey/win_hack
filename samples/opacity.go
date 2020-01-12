package samples

import (
	"main/win"

	"github.com/gonutz/w32"
)

//TestOpacity - Change opacity of another window
func TestOpacity() {
	//Change Opacity of window
	hwnd := w32.FindWindow("", "Безымянный — Блокнот")
	alpha := int32(70)
	lwaAlpha := int32(0x2)
	style := w32.GetWindowLongPtr(hwnd, w32.GWL_EXSTYLE)
	w32.SetWindowLongPtr(
		hwnd, w32.GWL_EXSTYLE, uintptr(uint32(style)|w32.WS_EX_LAYERED))

	win.SetLayeredWindowAttributes(uintptr(hwnd), 0, alpha, lwaAlpha)
	if alpha == 255 {
		w32.SetWindowLongPtr(
			hwnd, w32.GWL_EXSTYLE, uintptr(uint32(style)&^w32.WS_EX_LAYERED))
	}
}
