package injector

//Process Hollowing technique
func injectHollow() {
	//1.CreateProcess with CREATE_SUSPENDED flag
	//2.NtUnmapViewOfSection hide section in address map
	//3.WriteProcessMemory
	//4.ResumeThread
}

//Process DoppelGanging technique
func injectDoppelGanging() {
	//1.CreateTransaction ntfs
	//hTrans := CreateTransaction(null, 0, 0, 0, 0, 0, description)
	//2.CreateFileTransacted create temp file
	/*hTransFile := CreateFileTransacted(
		path_to_dummy_file,
		GENERIC_WRITE | GENERIC_READ,
		0,
		null,
		OPEN_EXISTING,
		FILE_ATTRIBUTE_NORMAL,
		null,
		hTrans,
		null,
		null
	)
	GetFileSizeEx(hTransFile, &pf_size)
	dwf_size = pf_size.LowPart
	byte *buf = malloc(dwf_size)
	*/
	//3.NtCreateSection buffers for temp files
	/*NtCreateSection(
		&hsection_obj,
		SECTION_ALL_ACCESS,
		nul,
		0,
		PAGE_READONLY,
		SEC_IMAGE,
		hTrans_file
	)
	*/
	//4.Check PEB
	//5.NtCreateProcessEx -> ResumeThread
	/*
		NtCreateProcessEx(
			&h_proc,
			GENERIC_ALL,
			null,
			GetCurrentProcess(),
			PS_INHERIT_HANDLES,
			hSection_obj,
			null,
			null,
			false
		)
		RtlCreateProcessParametersEx(
			&proc_parameters,
			&victim_path,
			null,
			null,
			&victim_path,
			null,
			null,
			null,
			null,
			null,
			RTL_USER_PROC_PARAMS_NORMALIZED
		)
		VirtualAllocEx()
		WriteProcessMemory()
		NtQueryInformationProcess()
		WriteProcessMemory()
		VirtualQueryEx()
		GetMappedFileName()
		NtCreateThreadEx(
			&hThread,
			GENERIC_ALL,
			null,
			h_proc,
			(LPTHREAD_START_ROUTINE)ep_proc,
			null,
			FALSE,
			0,
			0,
			0,
			null
		)
		RollbackTransaction(hTrans)
	*/
}
