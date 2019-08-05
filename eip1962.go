package eip1962

/*
#cgo CXXFLAGS: -std=c++17
#cgo CXXFLAGS: -I./include
#include "wrapper.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

const maxOutputLen = 256 * 3 * 2

var (
	ErrInvalidMsgLen = errors.New("invalid data length, need >= bytes")
	ErrCallFailed    = errors.New("library call returned an error")
)

// Call calls the C++ implementation of the EIP
func Call(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrInvalidMsgLen
	}
	ilen := len(data)
	outputBytes := make([]byte, maxOutputLen)
	olen := uint32(0)
	errStringBytes := make([]byte, maxOutputLen)
	errStringLen := uint32(0)

	var (
		inputdata  = (*C.char)(unsafe.Pointer(&data[0]))
		inputlen   = (C.uint32_t)(ilen)
		outputdata = (*C.char)(unsafe.Pointer(&outputBytes[0]))
		outputlen  = (*C.uint32_t)(unsafe.Pointer(&olen))
		errdata    = (*C.char)(unsafe.Pointer(&errStringBytes[0]))
		errlen     = (*C.uint32_t)(unsafe.Pointer(&errStringLen))
	)

	result := C.run(inputdata, inputlen, outputdata, outputlen, errdata, errlen)
	if result == 0 {
		// parse error string
		return nil, ErrCallFailed
	}

	return outputBytes[:olen], nil
}

// EstimateGas calls C++ implementation for a gas estimte
func EstimateGas(data []byte) uint64 {
	return 1000000
}
