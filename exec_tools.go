package mttools

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"syscall"
)

// Executes cmd and returns its output or err
func ExecCmd(cmd_path string, args []string) (string, error) {
	cmd := exec.Command(cmd_path, args...)

	buffer, err := cmd.CombinedOutput()

	return string(buffer), err
}

// Executes cmd and prints its output
func ExecCmdPrint(cmd_path string, args []string) error {
	out, err := ExecCmd(cmd_path, args)

	fmt.Print(out)
	return err
}

// Executes cmd printing its output to stdout during execution
func ExecCmdWaitAndPrint(cmd_path string, args []string) (output string, err error) {
	cmd := exec.Command(cmd_path, args...)
	//fmt.Println("CMD: " + cmd.String())

	var stdout_pipe, stderr_pipe io.ReadCloser

	if stdout_pipe, err = cmd.StdoutPipe(); err != nil {
		return "", err
	}

	if stderr_pipe, err = cmd.StderrPipe(); err != nil {
		return "", err
	}

	if err = cmd.Start(); err != nil {
		return "", err
	}

	stdout_scanner := bufio.NewScanner(stdout_pipe)
	stderr_scanner := bufio.NewScanner(stderr_pipe)

	f := true
	var text string
	for f {
		stdout := stdout_scanner.Scan()
		stderr := stderr_scanner.Scan()

		if stdout {
			text = stdout_scanner.Text()
			fmt.Println(text)
			output += text + "\n"
		}

		if stderr {
			text = stderr_scanner.Text()
			fmt.Println(text)
			output += text + "\n"
		}

		f = stdout || stderr
	}

	return output, cmd.Wait()
}

// Hides console only under Windows)
//
// Thanks to SyncThing!
// https://github.com/syncthing/syncthing/blob/main/lib/osutil/hidden_windows.go
func HideConsole() {
	if !IsWindows() {
		return
	}

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

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}
