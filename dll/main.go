package dll

import (
	"syscall"
	"unsafe"

	"github.com/castaneai/hinako"
	"github.com/nanitefactory/winmb"
)

import "C"

//export Test
func Test() {
	winmb.MessageBoxPlain("export Test", "export Test")
}

var hook *hinako.Hook

// OnProcessAttach is an async callback (hook).
//export OnProcessAttach
func OnProcessAttach(
	hinstDLL unsafe.Pointer, // handle to DLL module
	fdwReason uint32, // reason for calling function
	lpReserved unsafe.Pointer, // reserved
) {
	winmb.MessageBoxPlain("OnProcessAttach", "OnProcessAttach")
	hook = ntquery_hook()
}

// OnProcessDetach is an async callback (hook).
//export OnProcessDetach
func OnProcessDetach() {
	winmb.MessageBoxPlain("OnProcessDetach", "OnProcessDetach")
	defer hook.Close()
}

const title = "TITLE"

var version = "undefined"

//export WSTRPtr
func WSTRPtr(str string) uintptr {
	ptr, _ := syscall.UTF16PtrFromString(str)
	return uintptr(unsafe.Pointer(ptr))
}

func main() {

}
