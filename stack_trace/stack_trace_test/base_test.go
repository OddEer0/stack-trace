package stacktracetest

import (
	"context"
	"encoding/json"
	"testing"

	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	t.Run("default initial testing", func(t *testing.T) {
		ctx := stacktrace.Init(context.Background(), nil)
		sTrace, ok := ctx.Value(stacktrace.Key).(*stacktrace.StackTrace)
		assert.True(t, ok)
		assert.Equal(t, sTrace.IsLock, int32(0))
		assert.Equal(t, len(sTrace.Stack), 0)
		assert.Equal(t, cap(sTrace.Stack), stacktrace.DefaultStackCap)
	})

	t.Run("", func(t *testing.T) {
		ctx := stacktrace.Init(context.Background(), nil)
		stacktrace.Add(ctx, "1")
		stacktrace.Add(ctx, "2")
		stacktrace.Add(ctx, "3")
		stack := stacktrace.GetStack(ctx)
		res, err := json.Marshal(stack)
		assert.Nil(t, err)
		assert.Equal(t, string(res), `[{"message":"1"},{"message":"2"},{"message":"3"}]`)
		stacktrace.Done(ctx)
		stack = stacktrace.GetStack(ctx)
		res, err = json.Marshal(stack)
		assert.Nil(t, err)
		assert.Equal(t, string(res), `[{"message":"1"},{"message":"2"}]`)
	})
}
