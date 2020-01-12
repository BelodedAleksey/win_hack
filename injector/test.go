package injector

import (
	"log"
	"strconv"
)

func main() {
	log.Printf("Loading debug privileges...")
	loadDebugPrivileges()
	log.Printf("Done!")

	// dllFile, exeFile := os.Args[1], os.Args[2]
	// injectExe(dllFile, exeFile)

	dllFile, pidString := "my.dll", "pid"
	pid, err := strconv.ParseInt(pidString, 10, 64)
	must(err)
	injectPID(dllFile, pid)
}
