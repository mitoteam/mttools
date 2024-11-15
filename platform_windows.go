package mttools

import "syscall"

func hideConsoleHelper() {
	//import WinAPI function
	getConsoleWindowF := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	showWindowF := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")

	if getConsoleWindowF.Find() == nil && showWindowF.Find() == nil {
		hwnd, _, _ := getConsoleWindowF.Call()
		if hwnd != 0 {
			showWindowF.Call(hwnd, 0 /* 0 = hide */)
		}
	}
}
