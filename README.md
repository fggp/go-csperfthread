Go Binding for the Csound helper class CsoundPerformanceThread
========

This wrapper is still very experimental. It has been tested only on Linux.
It needs a proper installation of Csound with header files in the include path in the csound directory
(e.g. csound/csound.h). libcsound64 and libcsnd6 have to be in the PATH.

You can install this package with `go get`:

  `go get github.com/fggp/go-csperfthread`

Or you can download a zip archive of the project using the 'Download ZIP' button on the right.
You'll get a zip file named 'go-csperfthread-master.zip'. Decompressing it you'll get a directory named 'go-csperfthread-master'.
Rename this directory to 'go-csperfthread' and move it to '$GOPATH/src/github/fggp'. Enter into
the '$GOPATH/src/github/fggp/go-csperfthread' directory. You can eventually adapt the #cgo directives
in csnd.go to your system. Finally install the package with `go install`.

This wrapper is intended to be used with a double build of Csound.

Go version 1.2 or higher is needed for the C++ support of go build
