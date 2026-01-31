# Simple WebDriver wrapper in Go

```go
session, err := wdgo.New(context.Background(), "http://localhost:9516", &capability.Capabilities{
	ChromeOptions: &capability.Chrome{
		Binary: "/usr/bin/brave-browser",
		// Args:   []string{"headless=new"},
	},
})
if err != nil {
	log.Fatal(err)
}
defer session.Quit(context.Background())

session.Navigate(context.Background(), "https://example.com/")
```
