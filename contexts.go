package wdgo

import (
	"context"
	"fmt"
)

type (
	WindowType string

	NewWindowResult struct {
		Handle string     `json:"handle"`
		Type   WindowType `json:"type"`
	}
)

const (
	WindowTypeTab    WindowType = "tab"
	WindowTypeWindow WindowType = "window"
)

func (session *Session) WindowHandle(ctx context.Context) (string, error) {
	data, err := session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/window/handle", session.id), nil)
	if err != nil {
		return "", err
	}
	return parseResponse[string](data)
}

func (session *Session) WindowHandles(ctx context.Context) ([]string, error) {
	data, err := session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/window/handles", session.id), nil)
	if err != nil {
		return nil, err
	}
	return parseResponse[[]string](data)
}

func (session *Session) NewWindow(ctx context.Context, hint WindowType) (NewWindowResult, error) {
	data, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/window/new", session.id), map[string]string{
		"type": string(hint),
	})
	if err != nil {
		return NewWindowResult{}, err
	}
	return parseResponse[NewWindowResult](data)
}

func (session *Session) SwitchToWindow(ctx context.Context, handle string) error {
	_, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/window", session.id), map[string]string{
		"handle": handle,
	})
	return err
}

func (session *Session) CloseWindow(ctx context.Context) ([]string, error) {
	data, err := session.doRequest(ctx, "DELETE", fmt.Sprintf("session/%s/window", session.id), nil)
	if err != nil {
		return nil, err
	}
	return parseResponse[[]string](data)
}
