//          Copyright 2020, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

// Package bgproc starts a process in "background",
// i.e. a from terminal detached process.
package bgproc

// Process holds the name of the program and its arguments.
// PID is set after calling Start and if syscall.ForkExec
// has been used (not available on Windows).
type Process struct {
	Name string
	Args []string
	PID  int
}

// New creates and returns a new instance of Process. The
// parameter name is the absolute path to the program to start.
func New(name string) *Process {
	proc := new(Process)
	proc.Name = name
	return proc
}

// AddArg adds an argument to start the program with.
func (proc *Process) AddArg(arg string) {
	proc.Args = append(proc.Args, arg)
}

// String returns the name of the program and its arguments.
func (proc *Process) String() string {
	str := proc.Name
	for _, arg := range proc.Args {
		str += " " + arg
	}
	return str
}
