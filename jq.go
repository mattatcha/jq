package jq

/*
#cgo LDFLAGS: -ljq -all-static
#include <jq.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	errInvalidProgram = errors.New("program invalid")
	errJqStateNil     = errors.New("jq state nil")
)

// Jq used to hold state
type Jq struct {
	state *C.jq_state
}

// Free will release all memory from Jq state
func (j *Jq) Free() {
	C.jq_teardown(&j.state)
}

// NewJq returns an initialized state with a compiled program
// program should be a valid jq program/filter
// see http://stedolan.github.io/jq/manual/#Basicfilters for more info
func NewJq(program string) (*Jq, error) {
	jq := &Jq{
		state: C.jq_init(),
	}

	if jq.state == nil {
		return nil, errJqStateNil
	}

	pgm := C.CString(program)
	defer C.free(unsafe.Pointer(pgm))

	// Compiles a program into jq_state.
	if ok := C.jq_compile(jq.state, pgm); ok != 1 {
		jq.Free()
		return nil, errInvalidProgram
	}

	return jq, nil
}
