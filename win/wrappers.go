package win

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/gonutz/w32"
	"github.com/mitchellh/go-ps"
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

//EnumWindows func w32 not working!!!
func EnumWindows(callback func(window w32.HWND) uintptr) bool {
	f := syscall.NewCallback(func(w, _ uintptr) uintptr {
		return callback(w32.HWND(w))
	})
	ret, _, _ := enumWindowsProc.Call(f, 0)
	return ret != 0
}

//CreatePatternBrush func
func CreatePatternBrush(hBitmap w32.HBITMAP) w32.HBRUSH {
	ret, _, _ := createPatternBrushProc.Call(uintptr(hBitmap))
	return w32.HBRUSH(ret)
}

//SetActiveWindow func
func SetActiveWindow(hwnd w32.HWND) w32.HWND {
	ret, _, _ := setActiveWindowProc.Call(uintptr(hwnd))
	return w32.HWND(ret)
}

//FindWindowEx func
func FindWindowEx(parent, child w32.HWND, className, windowName string) w32.HWND {
	ret, _, _ := findWindowExProc.Call(
		uintptr(parent),
		uintptr(child),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowName))),
	)
	return w32.HWND(ret)
}

type window struct {
	Title  string
	Handle w32.HWND
	PID    int
}

func listAllWindows() (wins []*window, err error) {
	cb := func(hwnd w32.HWND) uintptr {
		if !w32.IsWindow(hwnd) || !w32.IsWindowVisible(hwnd) {
			return 1
		}
		title := ""
		tlen := w32.GetWindowTextLength(hwnd)
		if tlen != 0 {
			tlen++
			title = w32.GetWindowText(hwnd)
		}

		_, processID := w32.GetWindowThreadProcessId(hwnd)

		win := &window{
			Title:  title,
			Handle: hwnd,
			PID:    int(processID),
		}
		wins = append(wins, win)
		return 1
	}
	if !EnumWindows(cb) {
		return nil, fmt.Errorf("EnumWindows returned FALSE")
	}
	return wins, nil
}

func ancestors() []int {
	curr := os.Getpid()
	an := []int{curr}

	for {
		p, err := ps.FindProcess(curr)
		if p == nil || err != nil {
			break
		}
		curr = p.PPid()
		an = append(an, curr)
	}
	return an
}

func findFirstTarget(title string, wins []*window, ancestors []int) *window {
	if title == "" {
		for _, p := range ancestors {
			for _, w := range wins {
				if w.PID == p {
					return w
				}
			}
		}
	} else {
		t := strings.ToLower(title)

		for _, w := range wins {
			ancestor := false
			for _, p := range ancestors {
				if w.PID == p {
					ancestor = true
					break
				}
			}

			if t != "" && !ancestor {
				wt := strings.ToLower(w.Title)

				if strings.Contains(wt, t) {
					return w
				}
			} else if t == "" && ancestor {
				return w
			}
		}
	}
	return nil
}

//GetProcessHWND func
func GetProcessHWND(title string) (w32.HWND, error) {
	wins, err := listAllWindows()
	if err != nil {
		return 0, err
	}

	win := findFirstTarget(title, wins, ancestors())
	if win == nil {
		return 0, errors.New("no target")
	}
	return win.Handle, nil
}
