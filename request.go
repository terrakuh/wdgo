package wdgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (session *Session) get(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", session.endpoint+path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := session.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRemote, err)
	}
	defer resp.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read: %w", err)
	}

	return data, nil
}

func (session *Session) post(ctx context.Context, path string, payload any) ([]byte, error) {
	return session.doRequest(ctx, "POST", path, payload)
}

func (session *Session) doRequest(ctx context.Context, method, path string, payload any) ([]byte, error) {
	var body io.Reader

	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("marshal request: %w", err)
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, session.endpoint+path, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := session.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRemote, err)
	}
	defer resp.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read: %w", err)
	}

	return data, nil
}

func parseResponse[T any](data []byte) (T, error) {
	var empty T

	if err := parseAsError(data); err != nil {
		return empty, err
	}

	var resp struct {
		Value T `json:"value"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return empty, fmt.Errorf("bad response: %w", err)
	}
	return resp.Value, nil
}

func parseAsError(data []byte) error {
	var resp struct {
		Value *Error `json:"value"`
	}

	r := json.NewDecoder(bytes.NewReader(data))
	r.DisallowUnknownFields()

	if r.Decode(&resp) == nil && resp.Value != nil {
		return resp.Value
	}
	return nil
}
