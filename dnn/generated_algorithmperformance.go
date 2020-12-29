package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
import "runtime"

// AlgorithmPerformance is a representation of cudnnAlgorithmPerformance_t.
type AlgorithmPerformance struct {
	internal C.cudnnAlgorithmPerformance_t

	algoDesc *AlgorithmDescriptor
	status   Status
	time     float32
	memory   uintptr
}

// NewAlgorithmPerformance creates a new AlgorithmPerformance.
func NewAlgorithmPerformance(algoDesc *AlgorithmDescriptor, status Status, time float32, memory uintptr) (retVal *AlgorithmPerformance, err error) {
	var internal C.cudnnAlgorithmPerformance_t
	if err := result(C.cudnnCreateAlgorithmPerformance(&internal)); err != nil {
		return nil, err
	}

	if err := result(C.cudnnSetAlgorithmPerformance(internal, algoDesc.internal, status.C(), C.float(time), C.size_t(memory))); err != nil {
		return nil, err
	}

	retVal = &AlgorithmPerformance{
		internal: internal,
		algoDesc: algoDesc,
		status:   status,
		time:     time,
		memory:   memory,
	}
	runtime.SetFinalizer(retVal, destroyAlgorithmPerformance)
	return retVal, nil
}

// C returns the internal cgo representation.
func (a *AlgorithmPerformance) C() C.cudnnAlgorithmPerformance_t { return a.algoPerf }

// AlgoDesc returns the internal algoDesc.
func (a *AlgorithmPerformance) AlgoDesc() *AlgorithmDescriptor { return a.algoDesc }

// Status returns the internal status.
func (a *AlgorithmPerformance) Status() Status { return a.status }

// Time returns the internal time.
func (a *AlgorithmPerformance) Time() float32 { return a.time }

// Memory returns the internal memory.
func (a *AlgorithmPerformance) Memory() uintptr { return a.memory }

func destroyAlgorithmPerformance(obj *AlgorithmPerformance) {
	C.cudnnDestroyAlgorithmPerformance(obj.internal)
}
