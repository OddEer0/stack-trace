package stacktrace

import (
	"context"
	"sync"
	"sync/atomic"
)

func optDefault(opt *Option) *Option {
	if opt == nil {
		return &Option{
			Capacity: DefaultStackCap,
			IsLock:   false,
			ParceFn:  ParceFunc,
		}
	}

	if opt.Capacity == 0 {
		opt.Capacity = DefaultStackCap
	}

	if opt.ParceFn == nil {
		opt.ParceFn = ParceFunc
	}

	return opt
}

func Init(ctx context.Context, opt *Option) context.Context {
	opt = optDefault(opt)

	stackTrace := &StackTrace{
		Stack:   make([]any, 0, opt.Capacity),
		Mu:      sync.Mutex{},
		ParceFn: opt.ParceFn,
	}

	if opt.IsLock {
		stackTrace.IsLock = 1
	}

	return context.WithValue(ctx, Key, stackTrace)
}

func Add(ctx context.Context, trace any) {
	sTrace, ok := ctx.Value(Key).(*StackTrace)
	if ok {
		sTrace.Mu.Lock()
		defer sTrace.Mu.Unlock()
		sTrace.Stack = append(sTrace.Stack, sTrace.ParceFn(trace))
	}
}

func Done(ctx context.Context) {
	sTrace, ok := ctx.Value(Key).(*StackTrace)
	if ok {
		sTrace.Mu.Lock()
		defer sTrace.Mu.Unlock()
		sTrace.Stack = sTrace.Stack[:len(sTrace.Stack)-1]
	}
}

func GetStack(ctx context.Context) []any {
	sTrace, ok := ctx.Value(Key).(*StackTrace)
	if ok {
		sTrace.Mu.Lock()
		defer sTrace.Mu.Unlock()
		return sTrace.Stack
	}
	return nil
}

func Lock(ctx context.Context) {
	sTrace, ok := ctx.Value(Key).(*StackTrace)
	if ok && sTrace.IsLock == 0 {
		atomic.AddInt32(&sTrace.IsLock, 1)
	}
}

func Unlock(ctx context.Context) {
	sTrace, ok := ctx.Value(Key).(*StackTrace)
	if ok && sTrace.IsLock == 1 {
		atomic.AddInt32(&sTrace.IsLock, -1)
	}
}

func IsLock(ctx context.Context) bool {
	sTrace, ok := ctx.Value(Key).(*StackTrace)
	if ok && sTrace.IsLock == 1 {
		return true
	}
	return false
}
