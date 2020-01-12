package samples

import (
	"fmt"
	"main/win"
	"syscall"
	"unsafe"

	"github.com/gonutz/w32"
)

//TestMenu - func
func TestMenu() {
	hwnd := w32.FindWindow("", "Безымянный — Блокнот")
	text := w32.GetWindowText(hwnd)
	fmt.Println("Text: ", text)
	w32.SetWindowText(hwnd, "Замена")

	hmenu := w32.GetMenu(hwnd)
	numItems := w32.GetMenuItemCount(hmenu)

	for i := 0; i < numItems; i++ {
		str := w32.GetMenuString(hmenu, uint(i), w32.MF_BYPOSITION)
		fmt.Println("Меню: ", str)
		newStr := "Меню"
		var info w32.MENUITEMINFO
		info.Size = uint32(unsafe.Sizeof(info))
		info.Mask = win.MIIM_STRING
		info.TypeData = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(newStr)))
		w32.SetMenuItemInfo(hmenu, uint(i), true, &info)
		w32.DrawMenuBar(hwnd)
		id := w32.GetMenuItemID(hmenu, i)
		fmt.Println("ID: ", int32(id))
		if int32(id) == -1 {
			subMenu := w32.GetSubMenu(hmenu, i)
			numSubmenuItems := w32.GetMenuItemCount(subMenu)
			for j := 0; j < numSubmenuItems; j++ {
				substr := w32.GetMenuString(subMenu, uint(j), w32.MF_BYPOSITION)
				fmt.Println("Подменю: ", substr)
				info.TypeData = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Подменю")))
				w32.SetMenuItemInfo(subMenu, uint(j), true, &info)
				w32.DrawMenuBar(hwnd)
			}
		}
	}

}
