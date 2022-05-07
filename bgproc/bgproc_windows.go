//          Copyright 2020, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

package bgproc

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Start starts the program as a from terminal detached process.
func (proc *Process) Start() error {
	scriptPath, err := proc.createVBSkript()

	if err == nil {
		osCmd := exec.Command("wscript", scriptPath)
		err = osCmd.Start()

		if fileExists(scriptPath) {
			// wait 1.5 seconds until script runs
			time.Sleep(1500000000)
			os.Remove(scriptPath)
		}
	}
	return err
}

func (proc *Process) createVBSkript() (string, error) {
	script := "Set WshShell = CreateObject(\"WScript.Shell\")\r\n"
	script += "WshShell.Run \"\"\"" + proc.Name + "\"\""

	for _, arg := range proc.Args {
		script += " \"\"" + arg + "\"\""
	}
	script += "\", 0\r\n"
	script += "Set WshShell = Nothing\r\n"

	skriptPath := filepath.Join(os.TempDir(), "start.proc.vbs")
	skriptFile, err := os.OpenFile(skriptPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err == nil {
		defer skriptFile.Close()
		_, err = skriptFile.Write([]byte(script))
	}

	return skriptPath, err
}

func fileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || !os.IsNotExist(err)) && fileInfo != nil && !fileInfo.IsDir()
}
