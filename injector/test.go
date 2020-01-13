package injector

import (
	"log"
	"strconv"
)

//TestInject func
func TestInject() {
	log.Printf("Loading debug privileges...")
	loadDebugPrivileges()
	log.Printf("Done!")

	// dllFile, exeFile := os.Args[1], os.Args[2]
	// injectExe(dllFile, exeFile)

	dllFile, pidString := "E:\\GOPROJECTS\\win_hack\\dll\\my.dll", "18424"
	pid, err := strconv.ParseInt(pidString, 10, 64)
	must(err)
	injectPID(dllFile, pid)
}
