package main

import (
	"fmt"
	"main/win"
	"syscall"
	"unsafe"

	"github.com/castaneai/hinako"
	"github.com/zetamatta/go-outputdebug"
)

func ntquery_hook() *hinako.Hook {
	// API Hooking by hinako
	arch, err := hinako.NewRuntimeArch()
	if err != nil {
		outputdebug.String(fmt.Sprintf("NewRunTimeArch failed: %s", err.Error()))
	}
	var originalNtQuerySystemInformation *syscall.Proc = nil
	/*hook, err = hinako.NewHookByName(arch, "user32.dll", "MessageBoxW", func(hWnd syscall.Handle, lpText, lpCaption *uint16, uType uint) int {
		r, _, _ := originalMessageBoxW.Call(uintptr(hWnd), WSTRPtr("Hooked!"), WSTRPtr("Hooked!"), uintptr(uType))
		return int(r)
	})*/
	hook, err := hinako.NewHookByName(arch, "ntdll.dll", "NtQuerySystemInformation", func(SystemInformationClass uint32, SystemInformation uintptr, SystemInformationLength uint32, ReturnLength *uint32) int {
		//winmb.MessageBoxPlain("HOOK!", "HOOK!")

		// Make maxResults large for safety.
		// We can't invoke the api call with a results array that's too small.
		// If we have more than 2056 cores on a single host, then it's probably the future.
		maxBuffer := 2056
		// buffer for results from the windows proc
		resultBuffer := make([]win.SystemProcessInformation, maxBuffer)
		// size of the buffer in memory
		bufferSize := uintptr(win.SystemProcessInfoSize) * uintptr(maxBuffer)
		// size of the returned response
		var retSize uint32
		retCode, _, err := originalNtQuerySystemInformation.Call(
			win.System_Process_Information,            // System Information Class
			uintptr(unsafe.Pointer(&resultBuffer[0])), // pointer to first element in result buffer
			bufferSize,                        // size of the buffer in memory
			uintptr(unsafe.Pointer(&retSize)), // pointer to the size of the returned results the windows proc will set this
		)
		if retCode != 0 {
			outputdebug.String(fmt.Sprintf("NtQuerySystemInformation failed: %s , RetCode: %d", err.Error(), int(retCode)))
		}
		return int(retCode)
	})
	if err != nil {
		outputdebug.String(fmt.Sprintf("hook failed: %s", err.Error()))
	}
	originalNtQuerySystemInformation = hook.OriginalProc
	return hook
}
