package wdgo

import (
	"context"
	"fmt"
)

type (
	LocatorStrategy string
)

const (
	CSSSelector     LocatorStrategy = "css selector"
	LinkText        LocatorStrategy = "link text"
	PartialLinkText LocatorStrategy = "partial link text"
	TagName         LocatorStrategy = "tag name"
	XPath           LocatorStrategy = "xpath"
)

func (session *Session) FindElement(ctx context.Context, selector string, strategy LocatorStrategy) (*Element, error) {
	data, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/element", session.id), map[string]string{
		"using": string(strategy),
		"value": selector,
	})
	if err != nil {
		return nil, err
	}
	m, err := parseResponse[map[string]string](data)
	if err != nil {
		return nil, err
	} else if len(m) != 1 {
		return nil, fmt.Errorf("expected map of 1 but got: %d", len(m))
	}
	var id string
	for key := range m {
		id = key
	}
	return &Element{
		id:      id,
		session: session,
	}, nil
}

// func (session *Session) FindElements(ctx context.Context, selector string, strategy LocatorStrategy) ([]*Element, error) {
// 	data, err := session.doRequest(ctx, "POST", fmt.Sprintf("session/%s/element", session.id), map[string]string{
// 		"using": string(strategy),
// 		"value": selector,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	id, err := parseResponse[string](data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Element{
// 		id:      id,
// 		session: session,
// 	}, nil
// }
