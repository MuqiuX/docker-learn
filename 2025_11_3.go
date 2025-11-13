package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/moby/sys/reexec"
)

func main() {


	err := syscall.Unshare(syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = syscall.Sethostname([]byte("container"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	cmd := reexec.Command("/bin/bash")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
