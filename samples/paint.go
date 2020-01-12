package samples

import (
	"github.com/gonutz/w32"
)

//TestPaint - func
func TestPaint() {
	//Painting by hdc but it redraw by system
	hwnd := w32.FindWindow("", "Безымянный — Блокнот")
	prev := w32.POINT{}
	hdc := w32.GetDC(hwnd)
	w32.MoveToEx(hdc, 20, 20, &prev)
	w32.LineTo(hdc, 50, 50)
	w32.MoveToEx(hdc, int(prev.X), int(prev.Y), &prev)
	w32.ReleaseDC(hwnd, hdc)
}
