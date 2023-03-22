package mttools

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
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
func ExecCmdWaitPrint(cmd_path string, args []string) error {
	cmd := exec.Command(cmd_path, args...)
	//fmt.Println("CMD: " + cmd.String())

	var pipe io.ReadCloser
	var err error

	if pipe, err = cmd.StdoutPipe(); err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	return cmd.Wait()
}
