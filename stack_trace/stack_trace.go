package stacktrace

import "sync"

const (
	Key             = "eer0_stack_trace_request_id"
	DefaultStackCap = 10
)

type (
	StackTrace struct {
		Stack   []any
		IsLock  int32
		Mu      sync.Mutex
		ParceFn func(any) any
	}

	Option struct {
		Capacity uint
		IsLock   bool
		ParceFn  func(any) any
	}

	Message struct {
		Message string `json:"message"`
	}

	Method struct {
		Package string `json:"package"`
		Type    string `json:"type"`
		Method  string `json:"method"`
	}

	Func struct {
		Package  string `json:"package"`
		Function string `json:"function"`
	}
)
