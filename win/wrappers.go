package win

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/gonutz/w32"
)

//SetLayeredWindowAttributes func
func SetLayeredWindowAttributes(hwnd uintptr, crKey, bAlpha, dwFlags int32) bool {
	ret, _, _ := setLayeredWindowAttributesProc.Call(
		hwnd,
		uintptr(crKey),
		uintptr(bAlpha),
		uintptr(dwFlags),
	)
	return ret != 0
}

//EnumWindowsByTitle func
func EnumWindowsByTitle(title string) []w32.HWND {
	desktopHWND := w32.GetDesktopWindow()
	hwnd := w32.GetWindow(desktopHWND, w32.GW_CHILD)
	var arr = make([]w32.HWND, 0)
	for hwnd != 0 {
		str := w32.GetWindowText(hwnd)
		if strings.Contains(str, title) {
			fmt.Println(str)
			arr = append(arr, hwnd)
		}
		hwnd = w32.GetWindow(hwnd, w32.GW_HWNDNEXT)
	}
	return arr
}

// ToUnicode ...
func ToUnicode(wVirtKey uintptr, wScanCode uintptr, lpKeyState *[256]byte, pwszBuff *[256]byte, cchBuff int, wFlags uint) int {
	ret, _, _ := procToUnicode.Call(
		uintptr(wVirtKey),
		uintptr(wScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(pwszBuff)),
		uintptr(cchBuff),
		uintptr(wFlags))
	return int(ret)
}

// GetKeyboardState ...
func GetKeyboardState(lpKeyState *[256]byte) int {
	ret, _, _ := procGetKeyboardState.Call(uintptr(unsafe.Pointer(lpKeyState)))
	return int(ret)
}

// GetKeyboardLayoutName ...
func GetKeyboardLayoutName(pwszKLID *[256]byte) int {
	ret, _, _ := procGetKeyboardLayoutName.Call(uintptr(unsafe.Pointer(pwszKLID)))
	return int(ret)
}

// GetKeyboardLayout ...
func GetKeyboardLayout(idThread uintptr) int {
	ret, _, _ := procGetKeyboardLayout.Call(uintptr(idThread))
	return int(ret)
}
