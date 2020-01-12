package samples

import (
	"fmt"

	"github.com/StackExchange/wmi"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//GetSession func
func GetSession() (int, error) {
	type query struct {
		LogonId uint32
	}
	var s []query
	err := wmi.Query("Select LogonId from Win32_LogonSession", &s) //Select LogonId from Win32_LogonSession
	if err != nil {
		return 0, err
	}
	return len(s), nil
}

//TestLogon func
func TestLogon() {
	/*id, err := GetSession()
	if err != nil {
		fmt.Println("Error GetSession: ", err)
	}
	fmt.Println("ID: ", id)*/
	for {
		// init COM
		ole.CoInitialize(0)
		defer ole.CoUninitialize()

		unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
		if err != nil {
			fmt.Println("Error CreateObject: ", err)
		}
		defer unknown.Release()

		wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
		if err != nil {
			fmt.Println("Error QueryInterface: ", err)
		}
		defer wmi.Release()

		// service is a SWbemServices
		serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
		if err != nil {
			fmt.Println("Error Connect Server: ", err)
		}
		service := serviceRaw.ToIDispatch()
		defer service.Release()

		// result is a SWBemObjectSet
		resultRaw, err := oleutil.CallMethod(service, "ExecNotificationQuery",
			"Select * From __InstanceOperationEvent Within 5 Where TargetInstance Isa 'Win32_LogonSession'")
		if err != nil {
			fmt.Println("Error ExecQuery: ", err)
		}
		result := resultRaw.ToIDispatch()
		defer result.Release()

		// item is a SWbemObject, but really a Win32_Process
		itemRaw, err := oleutil.CallMethod(result, "NextEvent")
		if err != nil {
			fmt.Println("Error Next Event: ", err)
		}
		item := itemRaw.ToIDispatch()
		fmt.Println("LOGGGED ON!!!")
		objectRaw, err := oleutil.GetProperty(item, "TargetInstance")
		if err != nil {
			fmt.Println("Error GetProperty TargetInstance: ", err)
		}
		object := objectRaw.ToIDispatch()
		id, err := oleutil.GetProperty(object, "LogonId")
		if err != nil {
			fmt.Println("Error GetProperty LogonId : ", err)
		}
		fmt.Println("ID: ", id.ToString())
		defer item.Release()

	}
}
