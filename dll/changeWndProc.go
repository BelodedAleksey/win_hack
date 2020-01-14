package main

import (
	"fmt"
	"syscall"

	"github.com/gonutz/w32"
	"github.com/nanitefactory/winmb"
	"github.com/tadvi/winc"
	"github.com/zetamatta/go-outputdebug"
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
	var canvas *winc.Canvas
	//bitmap, _ := winc.NewBitmapFromFile("E:\\GOPROJECTS\\win_hack\\assets\\image.bmp", winc.RGB(255, 0, 255))
	wndProc := func(hwnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		outputdebug.String(fmt.Sprintf("Message: %s wParam: %s lParam: %s", msg, wParam, lParam))
		switch msg {
		case w32.WM_COMMAND:
			switch wParam {
			case 2000:
				winmb.MessageBoxPlain("BUTTON: ", "Хоббит, сэр!")
				break
			}
			break
		case w32.WM_MENUSELECT:
			//w32.LOWORD(uint32(wParam)) - id of menu item by position
			//lparam - handle of menu
			//Select menu item on 0 position
			/*if w32.LOWORD(uint32(wParam)) == 0 && lParam == uintptr(hmenu) {
				winmb.MessageBoxPlain("MENU: ", "File clicked!")
			}*/
			break
		case w32.WM_PAINT:
			//Call with any redraw
			//w32.InvalidateRect(hwnd, nil, true)
			//w32.UpdateWindow(hwnd)

			//hdc := w32.GetDC(hwnd)
			//defer w32.ReleaseDC(hwnd, hdc)
			//Draw line
			/*prev := w32.POINT{}
			w32.MoveToEx(hdc, 20, 20, &prev)
			w32.LineTo(hdc, 50, 50)
			w32.MoveToEx(hdc, int(prev.X), int(prev.Y), &prev)*/
			//Draw Text
			//w32.TextOut(hdc, 5, 5, "Проверочка!!!")

			//Draw Bitmap
			//canvas = winc.NewCanvasFromHwnd(winc32.HWND(hwnd))
			//canvas.DrawBitmap(bitmap, 0, 0)
			break
		case w32.WM_DESTROY:
			canvas.Dispose()
			break
		}
		//return w32.DefWindowProc(hwnd, msg, wParam, lParam)
		return w32.CallWindowProc(origProc, hwnd, msg, wParam, lParam)
	}
	w32.SetWindowLongPtr(
		hwnd, w32.GWLP_WNDPROC, syscall.NewCallback(wndProc),
	)
}
