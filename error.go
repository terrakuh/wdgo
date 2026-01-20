package wdgo

import "errors"

type (
	Error struct {
		Code       ErrCode        `json:"error"`
		Message    string         `json:"message"`
		Stacktrace string         `json:"stacktrace"`
		Data       map[string]any `json:"data"`
	}

	ErrCode string
)

const (
	ErrCodeElementClickIntercepted ErrCode = "element click intercepted"
	ErrCodeElementNotInteractable  ErrCode = "element not interactable"
	ErrCodeInsecureCertificate     ErrCode = "insecure certificate"
	ErrCodeInvalidArgument         ErrCode = "invalid argument"
	ErrCodeInvalidCookieDomain     ErrCode = "invalid cookie domain"
	ErrCodeInvalidElementState     ErrCode = "invalid element state"
	ErrCodeInvalidSelector         ErrCode = "invalid selector"
	ErrCodeInvalidSessionID        ErrCode = "invalid session id"
	ErrCodeJavascriptError         ErrCode = "javascript error"
	ErrCodeMoveTargetOutOfBounds   ErrCode = "move target out of bounds"
	ErrCodeNoSuchAlert             ErrCode = "no such alert"
	ErrCodeNoSuchCookie            ErrCode = "no such cookie"
	ErrCodeNoSuchElement           ErrCode = "no such element"
	ErrCodeNoSuchFrame             ErrCode = "no such frame"
	ErrCodeNoSuchWindow            ErrCode = "no such window"
	ErrCodeNoSuchShadowRoot        ErrCode = "no such shadow root"
	ErrCodeScriptTimeoutError      ErrCode = "script timeout"
	ErrCodeSessionNotCreated       ErrCode = "session not created"
	ErrCodeStaleElementReference   ErrCode = "stale element reference"
	ErrCodeDetachedShadowRoot      ErrCode = "detached shadow root"
	ErrCodeTimeout                 ErrCode = "timeout"
	ErrCodeUnableToSetCookie       ErrCode = "unable to set cookie"
	ErrCodeUnableToCaptureScreen   ErrCode = "unable to capture screen"
	ErrCodeUnexpectedAlertOpen     ErrCode = "unexpected alert open"
	ErrCodeUnknownCommand          ErrCode = "unknown command"
	ErrCodeUnknownError            ErrCode = "unknown error"
	ErrCodeUnknownMethod           ErrCode = "unknown method"
	ErrCodeUnsupportedOperation    ErrCode = "unsupported operation"
)

func (err Error) Error() string {
	return err.Message
}

func IsErrCode(err error, code ErrCode) bool {
	var e Error
	return errors.As(err, &e) && e.Code == code
}
