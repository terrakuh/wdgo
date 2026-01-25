package wdgo

import (
	"context"
	"fmt"
)

func (session *Session) CurrentURL(ctx context.Context) (string, error) {
	data, err := session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/url", session.id), nil)
	if err != nil {
		return "", err
	}
	return parseResponse[string](data)
}

func (session *Session) Title(ctx context.Context) (string, error) {
	data, err := session.doRequest(ctx, "GET", fmt.Sprintf("session/%s/title", session.id), nil)
	if err != nil {
		return "", err
	}
	return parseResponse[string](data)
}

func (session *Session) Navigate(ctx context.Context, url string) error {
	_, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/url", session.id), map[string]any{"url": url})
	return err
}

func (session *Session) Refresh(ctx context.Context) error {
	_, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/refresh", session.id), map[string]string{})
	return err
}
