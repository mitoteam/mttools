package mttools

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

// Executes cmd printing its stdout+stderr to current stdout during execution and returns its output after
func ExecCmdWaitAndPrint(cmd_path string, args []string) (output string, err error) {
	cmd := exec.Command(cmd_path, args...)
	//fmt.Println("CMD: " + cmd.String())

	var sb strings.Builder
	mw := io.MultiWriter(&sb, os.Stdout)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err = cmd.Start(); err != nil {
		return "", err
	}

	err = cmd.Wait()

	output = sb.String()

	return output, err
}

// Executes whole string in command-line shell. Returns its output and/or error.
func ExecCommandLine(command_line string) (string, error) {
	var cmd_path string
	var args []string

	if IsWindows() {
		cmd_path = "cmd.exe"
		args = []string{"/C", command_line}
	} else if IsLinux() {
		cmd_path = "sh"
		args = []string{"-c", command_line}
	} else {
		log.Panicln("Unknown platform in ExecCommandLine()")
	}

	return ExecCmd(cmd_path, args)
}

// Hides console (only under Windows)
//
// Thanks to SyncThing!
// https://github.com/syncthing/syncthing/blob/main/lib/osutil/hidden_windows.go
func HideConsole() {
	hideConsoleHelper()
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}
