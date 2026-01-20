package wdgo

import (
	"net/http"

	"github.com/terrakuh/wdgo/capability"
)

type Option func(ctx *setupContext)

func WithClient(client *http.Client) Option {
	return func(ctx *setupContext) {
		ctx.session.client = client
	}
}

func WithFirstMatch(match *capability.Capabilities) Option {
	return func(ctx *setupContext) {
		ctx.firstMatches = append(ctx.firstMatches, match)
	}
}
