package capability

type Capabilities struct {
	BrowserName    string `json:"browserName,omitempty"`
	BrowserVersion string `json:"browserVersion,omitempty"`

	ChromeOptions *Chrome `json:"goog:chromeOptions,omitempty"`
}
