package win

import (
	"unsafe"

	"github.com/gonutz/w32"
)

type SystemProcessorPerformanceInformation struct {
	IdleTime       int64 // idle time in 100ns (this is not a filetime).
	KernelTime     int64 // kernel time in 100ns.  kernel time includes idle time. (this is not a filetime).
	UserTime       int64 // usertime in 100ns (this is not a filetime).
	DpcTime        int64 // dpc time in 100ns (this is not a filetime).
	InterruptTime  int64 // interrupt time in 100ns
	InterruptCount uint32
}

// KPRIORITY
type kPriority int32

// UNICODE_STRING
type unicodeString struct {
	Length        uint16
	MaximumLength uint16
	BufferAddr    *uint16
}

// SYSTEM_PROCESS_INFORMATION
type SystemProcessInformation struct {
	NextEntryOffset              uint32        // ULONG
	NumberOfThreads              uint32        // ULONG
	WorkingSetPrivateSize        int64         // LARGE_INTEGER
	HardFaultCount               uint32        // ULONG
	NumberOfThreadsHighWatermark uint32        // ULONG
	CycleTime                    uint64        // ULONGLONG
	CreateTime                   int64         // LARGE_INTEGER
	UserTime                     int64         // LARGE_INTEGER
	KernelTime                   int64         // LARGE_INTEGER
	ImageName                    unicodeString // UNICODE_STRING
	BasePriority                 kPriority     // KPRIORITY
	UniqueProcessID              uintptr       // HANDLE
	InheritedFromUniqueProcessID uintptr       // HANDLE
	HandleCount                  uint32        // ULONG
	SessionID                    uint32        // ULONG
	UniqueProcessKey             *uint32       // ULONG_PTR
	PeakVirtualSize              uintptr       // SIZE_T
	VirtualSize                  uintptr       // SIZE_T
	PageFaultCount               uint32        // ULONG
	PeakWorkingSetSize           uintptr       // SIZE_T
	WorkingSetSize               uintptr       // SIZE_T
	QuotaPeakPagedPoolUsage      uintptr       // SIZE_T
	QuotaPagedPoolUsage          uintptr       // SIZE_T
	QuotaPeakNonPagedPoolUsage   uintptr       // SIZE_T
	QuotaNonPagedPoolUsage       uintptr       // SIZE_T
	PagefileUsage                uintptr       // SIZE_T
	PeakPagefileUsage            uintptr       // SIZE_T
	PrivatePageCount             uintptr       // SIZE_T
	ReadOperationCount           int64         // LARGE_INTEGER
	WriteOperationCount          int64         // LARGE_INTEGER
	OtherOperationCount          int64         // LARGE_INTEGER
	ReadTransferCount            int64         // LARGE_INTEGER
	WriteTransferCount           int64         // LARGE_INTEGER
	OtherTransferCount           int64         // LARGE_INTEGER
}

const (
	System_Basic_Information                     = 0
	System_Performance_Information               = 2
	System_TimeOfDay_Information                 = 3
	System_Process_Information                   = 5
	System_Processor_PerformanceInformation      = 8
	System_Interrupt_Information                 = 23
	System_Exception_Information                 = 33
	System_Registry_Quota_Information            = 37
	System_Lookaside_Information                 = 45
	System_Processor_Idle_Cycle_Time_Information = 83
	System_Processor_Cycle_Time_Information      = 108
	System_Policy_Information                    = 134

	// size of systemProcessInfoSize in memory
	SystemProcessInfoSize = unsafe.Sizeof(SystemProcessInformation{})
)

// MSLLHOOKSTRUCT ...
type MSLLHOOKSTRUCT struct {
	Pt          w32.POINT
	MouseData   uintptr
	Flags       uintptr
	Time        uintptr
	DwExtraInfo uintptr
}
