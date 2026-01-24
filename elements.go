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
	return session.findElement(ctx, fmt.Sprintf("session/%s/element", session.id), selector, strategy)
}

func (session *Session) FindElements(ctx context.Context, selector string, strategy LocatorStrategy) ([]*Element, error) {
	return session.findElements(ctx, fmt.Sprintf("session/%s/elements", session.id), selector, strategy)
}

func (session *Session) findElement(ctx context.Context, path, selector string, strategy LocatorStrategy) (*Element, error) {
	data, err := session.doRequest(ctx, "POST", path, map[string]string{
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

	for _, key := range m {
		return &Element{
			id:      key,
			session: session,
		}, nil
	}
	panic("unreachable")
}

func (session *Session) findElements(ctx context.Context, path, selector string, strategy LocatorStrategy) ([]*Element, error) {
	data, err := session.doRequest(ctx, "POST", path, map[string]string{
		"using": string(strategy),
		"value": selector,
	})
	if err != nil {
		return nil, err
	}
	maps, err := parseResponse[[]map[string]string](data)
	if err != nil {
		return nil, err
	}
	elements := make([]*Element, 0, len(maps))
	for _, m := range maps {
		if len(m) != 1 {
			return nil, fmt.Errorf("expected map of 1 but got: %d", len(m))
		}
		for _, key := range m {
			elements = append(elements, &Element{id: key, session: session})
		}
	}
	return elements, nil
}
