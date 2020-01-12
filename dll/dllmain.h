#include <windows.h>

void OnProcessAttach(HINSTANCE, DWORD, LPVOID);

void OnProcessDetach();

BOOL WINAPI DllMain(
    HINSTANCE _hinstDLL, // handle to DLL module
    DWORD _fdwReason,    // reason for calling function
    LPVOID _lpReserved   // reserved
);
