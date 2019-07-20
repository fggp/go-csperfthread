package csperfthread

/*
#cgo CXXFLAGS: -DUSE_DOUBLE=1

#include <csound/csound.h>
#include "csound_pt.h"
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -lcsound64 -lcsnd6
*/
import "C"

import (
	csnd "github.com/fggp/go-csnd"
	"unsafe"
)

type CsoundPerformanceThread struct {
	cpt (C.Cpt)
}

// Return a new CsoundPerformanceThread object.
func NewCsoundPerformanceThread(csound csnd.CSOUND) CsoundPerformanceThread {
	cpt := C.NewCsoundPT((*C.struct_CSOUND_)(csound.Cs))
	return CsoundPerformanceThread{cpt}
}

// Free the memory associated with the underlying C++ object.
func (pt CsoundPerformanceThread) Delete() {
	C.DeleteCsoundPT(pt.cpt)
	pt.cpt = nil
}

////////////////////////////////////////////////////////////////

// Workaround to avoid the 'cgo argument has Go pointer to Go pointer'
// runtime error (since go1.6)
var registry = make(map[int]unsafe.Pointer)
var index int

type PTprocessHandler func(cbData unsafe.Pointer)

var ptProcess PTprocessHandler

// Return the process callback as a PTprocessHandler.
func (pt CsoundPerformanceThread) ProcessCallback() PTprocessHandler {
	return ptProcess
}

//export goPTprocessCB
func goPTprocessCB(cbData unsafe.Pointer) {
	if ptProcess == nil {
		return
	}
	dataP := registry[*(*int)(cbData)]
	ptProcess(dataP)
}

// Set the process callback.
func (pt CsoundPerformanceThread) SetProcessCallback(f PTprocessHandler, cbData unsafe.Pointer) {
	ptProcess = f
	registry[index] = cbData
	C.CsoundPTsetProcessCB(pt.cpt, unsafe.Pointer(&index))
}

////////////////////////////////////////////////////////////////

// Tell if performance thread is running.
func (pt CsoundPerformanceThread) IsRunning() bool {
	return C.CsoundPTisRunning(pt.cpt) != 0
}

// Return the Csound instance pointer.
func (pt CsoundPerformanceThread) GetCsound() *C.CSOUND {
	return C.CsoundPTgetCsound(pt.cpt)
}

// Return the current status, zero if still playing, positive if
// the end of score was reached or performance was stopped, and
// negative if an error occured.
func (pt CsoundPerformanceThread) GetStatus() int {
	return int(C.CsoundPTgetStatus(pt.cpt))
}

// Continue performance if it was paused.
func (pt CsoundPerformanceThread) Play() {
	C.CsoundPTplay(pt.cpt)
}

// Pause performance (can be continued by calling Play()).
func (pt CsoundPerformanceThread) Pause() {
	C.CsoundPTpause(pt.cpt)
}

// Pause performance unless it is already paused, in which case
// it is continued.
func (pt CsoundPerformanceThread) TogglePause() {
	C.CsoundPTtogglePause(pt.cpt)
}

// Stop performance (cannot be continued).
func (pt CsoundPerformanceThread) Stop() {
	C.CsoundPTstop(pt.cpt)
}

// Starts recording the output from Csound. The sample rate and number
// of channels are taken directly from the running Csound instance.
// If args are specified, they are samplebits and numbufs respectively.
// If args are not specified, samplebits default to 16 and numbufs to 4.
func (pt CsoundPerformanceThread) Record(filename string, args ...int) {
	var cfname *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(cfname))
	samplebits, numbufs := 16, 4
	if len(args) >= 1 {
		samplebits = args[0]
		if len(args) == 2 {
			numbufs = args[1]
		}
	}
	C.CsoundPTrecord(pt.cpt, cfname, C.int(samplebits), C.int(numbufs))
}

// Stops recording and closes audio file.
func (pt CsoundPerformanceThread) StopRecord() {
	C.CsoundPTstopRecord(pt.cpt)
}

// Send a score event of type 'opcod' (e.g. 'i' for a note event), with
// p-fields in array 'p' (p[0] is p1). If absp2mode is true,
// the start time of the event is measured from the beginning of
// performance, instead of the default of relative to the current time.
func (pt CsoundPerformanceThread) ScoreEvent(absp2mode bool, opcod byte, p []csnd.MYFLT) {
	var absolute C.int
	if absp2mode {
		absolute = 1
	}
	C.CsoundPTscoreEvent(pt.cpt, absolute, C.char(opcod), C.int(len(p)),
		(*C.double)(&p[0]))
}

// Send a score event as a string, similarly to line events (-L).
func (pt CsoundPerformanceThread) InputMessage(s string) {
	var cmsg *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(cmsg))
	C.CsoundPTinputMessage(pt.cpt, cmsg)
}

// Set the playback time pointer to the specified value (in seconds).
func (pt CsoundPerformanceThread) SetScoreOffsetSeconds(timeVal float64) {
	C.CsoundPTsetScoreOffsetSeconds(pt.cpt, C.double(timeVal))
}

// Wait until the performance is finished or fails, and return a
// positive value if the end of score was reached or Stop() was called,
// and a negative value if an error occured. Also releases any resources
// associated with the performance thread object.
func (pt CsoundPerformanceThread) Join() int {
	return int(C.CsoundPTjoin(pt.cpt))
}

// Wait until all pending messages (pause, send score event, etc.)
// are actually received by the performance thread.
func (pt CsoundPerformanceThread) FlushMessageQueue() {
	C.CsoundPTflushMessageQueue(pt.cpt)
}
