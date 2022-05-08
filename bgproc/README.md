# bgproc

[![Go Reference](https://pkg.go.dev/badge/github.com/vbsw/golib/bgproc.svg)](https://pkg.go.dev/github.com/vbsw/golib/bgproc) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/golib/bgproc)](https://goreportcard.com/report/github.com/vbsw/golib/bgproc) [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

## About
bgproc is a package for Go to start a process in "background", i.e. as a from terminal detached process. (It is encouraged to use something else than this, like services on Windows or daemons on Linux.) bgproc is published on <https://github.com/vbsw/golib/bgproc> and <https://gitlab.com/vbsw/golib/bgproc>.

## Copyright
Copyright 2020, Vitali Baumtrok (vbsw@mailbox.org).

bgproc is distributed under the Boost Software License, version 1.0. (See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)

bgproc is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Example

	import (
		"fmt"
		"github.com/vbsw/golib/bgproc"
		"os"
		"path/filepath"
	)

	func main() {
		if len(os.Args) == 2 && os.Args[1] == "--background" {
			progName, err := filepath.Abs(os.Args[0])

			if err == nil {
				proc := bgproc.New(progName)
				err = proc.Start()
			}
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

## References
- https://go.dev/doc/install
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
- https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool
