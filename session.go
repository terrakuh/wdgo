package wdgo

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/terrakuh/wdgo/capability"
)

type (
	Session struct {
		endpoint string
		id       string
		client   *http.Client
	}

	setupContext struct {
		session      *Session
		firstMatches []*capability.Capabilities
	}

	newSessionResp struct {
		SessionID string `json:"sessionId"`
	}
)

var ErrRemote = errors.New("bad remote")

func New(ctx context.Context, endpoint string, capabilities *capability.Capabilities, options ...Option) (*Session, error) {
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}
	session := &Session{
		endpoint: endpoint,
		client:   http.DefaultClient,
	}
	setupCtx := setupContext{session: session}
	for _, opt := range options {
		opt(&setupCtx)
	}

	data, err := session.doRequest(ctx, "POST", "session", map[string]any{"capabilities": map[string]any{
		"alwaysMatch":  capabilities,
		"firstMatches": setupCtx.firstMatches,
	}})
	if err != nil {
		return nil, err
	}

	resp, err := parseResponse[newSessionResp](data)
	if err != nil {
		return nil, err
	}
	session.id = resp.SessionID

	return session, nil
}

func (session *Session) Delete(ctx context.Context) error {
	_, err := session.doRequest(ctx, "DELETE", "session/"+session.id, nil)
	return err
}

// Close closes all open windows and deletes the session.
func (session *Session) Close(ctx context.Context) error {
	for {
		remaining, err := session.CloseWindow(ctx)
		if err != nil {
			return err
		} else if len(remaining) == 0 {
			break
		}
		if err = session.SwitchToWindow(ctx, remaining[0]); err != nil {
			return err
		}
	}
	return session.Delete(ctx)
}
