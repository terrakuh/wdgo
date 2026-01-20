package capability

type (
	Chrome struct {
		/// List of command-line arguments to use when starting Chrome. Arguments with an associated value should be
		/// separated by a '=' sign (e.g., ['start-maximized', 'user-data-dir=/tmp/temp_profile']). See
		/// [here](http://peter.sh/experiments/chromium-command-line-switches/) for a list of Chrome arguments.
		Args []string `json:"args,omitempty"`
		/// Path to the Chrome executable to use (on Mac OS X, this should be the actual binary, not just the app.
		/// e.g., '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome')
		Binary string `json:"binary,omitempty"`
		/// A list of Chrome extensions to install on startup. Each item in the list should be a base-64 encoded
		/// packed Chrome extension (.crx)
		Extensions []string `json:"extensions,omitempty"`
		/// A dictionary with each entry consisting of the name of the preference and its value. These preferences
		/// are applied to the Local State file in the user data folder.
		LocalState map[string]any `json:"localState,omitempty"`
		/// A dictionary with each entry consisting of the name of the preference and its value. These preferences
		/// are only applied to the user profile in use. See the 'Preferences' file in Chrome's user data directory
		/// for examples.
		Preferences map[string]any `json:"preferences,omitempty"`
		/// If false, Chrome will be quit when ChromeDriver is killed, regardless of whether the session is quit. If
		/// true, Chrome will only be quit if the session is quit (or closed). Note, if true, and the session is not
		/// quit, ChromeDriver cannot clean up the temporary user data directory that the running Chrome instance is
		/// using.
		Detach bool `json:"detach"`
		/// An address of a Chrome debugger server to connect to, in the form of <hostname/ip:port>, e.g.
		/// '127.0.0.1:38947'
		DebuggerAddress string `json:"debuggerAddress,omitempty"`
		/// List of Chrome command line switches to exclude that ChromeDriver by default passes when starting
		/// Chrome. Do not prefix switches with --.
		ExcludeSwitches []string `json:"excludeSwitches,omitempty"`
		/// Directory to store Chrome minidumps (Supported only on Linux.)
		MinidumpPath string `json:"minidumpPath,omitempty"`
		/// A dictionary with either a value for “deviceName,” or values for “deviceMetrics” and “userAgent.” Refer
		/// to Mobile Emulation for more information.
		MobileEmulation map[string]any `json:"mobileEmulation,omitempty"`
		/// An optional dictionary that specifies performance logging preferences.
		PerformanceLoggingPreferences *ChromePerformanceLoggingPreferences `json:"perfLoggingPrefs,omitempty"`
		/// A list of window types that will appear in the list of window handles. For access to <webview> elements,
		/// include "webview" in this list.
		WindowTypes []string `json:"windowTypes,omitempty"`
	}

	ChromePerformanceLoggingPreferences struct {
		/// Whether or not to collect events from Network domain.
		EnableNetwork bool `json:"enableNetwork,omitempty"`
		/// Whether or not to collect events from Page domain.
		EnablePage bool `json:"enablePage,omitempty"`
		/// A comma-separated string of Chrome tracing categories for which trace events should be collected. An
		/// unspecified or empty string disables tracing.
		TraceCategories string `json:"traceCategories,omitempty"`
		/// The requested number of milliseconds between DevTools trace buffer usage events. For example, if 1000,
		/// then once per second, DevTools will report how full the trace buffer is. If a report indicates the
		/// buffer usage is 100%, a warning will be issued.
		BufferUsageReportingInterval int `json:"bufferUsageReportingInterval"`
	}
)
