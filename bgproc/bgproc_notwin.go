//          Copyright 2020, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

//go:build !windows
// +build !windows

package bgproc

import (
	"os"
	"syscall"
)

// Start starts the program as a from terminal detached process.
func (proc *Process) Start() error {
	args := proc.argsForFork()
	attr := syscall.ProcAttr{
		"",
		[]string{},
		[]uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		nil,
	}
	pid, err := syscall.ForkExec(proc.Name, args, &attr)

	if err == nil {
		proc.PID = pid
	} else {
		proc.PID = 0
	}
	return err
}

func (proc *Process) argsForFork() []string {
	args := make([]string, 1, len(proc.Args)+1)
	args[0] = proc.Name
	args = append(args, proc.Args...)
	return args
}
