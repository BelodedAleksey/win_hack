package main

import (
	"main/win"
	"syscall"

	"github.com/gonutz/w32"
	"github.com/nanitefactory/winmb"
)

//changeWndProc func
func changeWndProc() {
	hwnd := w32.FindWindow("", "Безымянный — Блокнот")
	//hwnd := win.GetProcessHWND("Блокнот") CRUSHED PROGRAMM!!!
	hmenu := w32.GetMenu(hwnd)
	newMenu := w32.CreateMenu()
	w32.AppendMenu(hmenu, w32.MF_STRING|w32.MF_POPUP, uintptr(newMenu), "Кто?")
	w32.AppendMenu(newMenu, w32.MF_STRING, 2000, "Button")
	w32.DrawMenuBar(hwnd)
	//Painting by catch WM_PAINT
	origProc := w32.GetWindowLongPtr(hwnd, w32.GWLP_WNDPROC)
	wndProc := func(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		switch msg {
		case w32.WM_COMMAND:
			switch wParam {
			case 2000:
				winmb.MessageBoxPlain("BUTTON: ", "Хоббит, сэр!")
				break
			}
			break
		case w32.WM_MENUSELECT:
			//win.LOWORD(uint32(wParam)) - id of menu item by position
			//lparam - handle of menu
			if win.LOWORD(uint32(wParam)) == 0 && lParam == uintptr(hmenu) {
				winmb.MessageBoxPlain("MENU: ", "File clicked!")
			}
			break
		case w32.WM_PAINT:
			w32.InvalidateRect(hwnd, nil, true)
			w32.UpdateWindow(hwnd)
			prev := w32.POINT{}
			hdc := w32.GetDC(hwnd)
			w32.MoveToEx(hdc, 20, 20, &prev)
			w32.LineTo(hdc, 50, 50)
			w32.MoveToEx(hdc, int(prev.X), int(prev.Y), &prev)
			w32.ReleaseDC(hwnd, hdc)
			break
		}
		//return w32.DefWindowProc(hwnd, msg, wParam, lParam)
		return w32.CallWindowProc(origProc, hwnd, msg, wParam, lParam)
	}
	w32.SetWindowLongPtr(
		hwnd, w32.GWLP_WNDPROC, syscall.NewCallback(wndProc),
	)
}
