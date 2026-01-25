package wdgo

import (
	"context"
	"fmt"
	"math"
	"time"
)

func (session *Session) ExecuteScript(ctx context.Context, script string, args []any, timeout time.Duration) (any, error) {
	var timeoutVal *int32
	if timeout > 0 {
		v, err := safeInt64To32(timeout.Milliseconds())
		if err != nil {
			return nil, err
		}
		timeoutVal = &v
	}
	if args == nil {
		args = []any{}
	}
	data, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/execute/sync", session.id), map[string]any{
		"script":  script,
		"timeout": timeoutVal,
		"args":    args,
	})
	if err != nil {
		return nil, err
	}
	return parseResponse[any](data)
}

func safeInt64To32(value int64) (int32, error) {
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("%w: value out of int32 range", ErrInvalidArg)
	}
	return int32(value), nil
}
