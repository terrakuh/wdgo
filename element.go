package wdgo

import (
	"context"
	"encoding/base64"
	"fmt"
)

type Element struct {
	id      string
	session *Session
}

func (element *Element) Text(ctx context.Context) (string, error) {
	data, err := element.session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/element/%s/text", element.session.id, element.id), nil)
	if err != nil {
		return "", err
	}
	return parseResponse[string](data)
}

func (element *Element) Property(ctx context.Context, name string) (string, error) {
	data, err := element.session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/element/%s/property/%s", element.session.id, element.id, name), nil)
	if err != nil {
		return "", err
	}
	return parseResponse[string](data)
}

func (element *Element) Screenshot(ctx context.Context, text string) ([]byte, error) {
	data, err := element.session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/element/%s/screenshot", element.session.id, element.id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := parseResponse[string](data)
	if err != nil {
		return nil, err
	}
	data, err = base64.StdEncoding.DecodeString(resp)
	if err != nil {
		return nil, fmt.Errorf("bad base64: %w", err)
	}
	return data, nil
}

func (element *Element) Click(ctx context.Context) error {
	_, err := element.session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/element/%s/click", element.session.id, element.id), nil)
	return err
}

func (element *Element) Clear(ctx context.Context) error {
	_, err := element.session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/element/%s/clear", element.session.id, element.id), nil)
	return err
}

func (element *Element) SendKeys(ctx context.Context, text string) error {
	_, err := element.session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/element/%s/value", element.session.id, element.id), map[string]string{"text": text})
	return err
}
