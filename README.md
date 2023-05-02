# uk-crimes
Simple API client for UK Crimes at particular coordinates

```
go get -u github.com/qazaqpyn/uk-crimes
```

```go
// New return UKCrime client
func New() (*Client, error)

// Client is custom UKCrime client containing api link
type Client struct
    // ScanCoordinates send coordinates for checking and gets report back
    ScanCoordinates(longitude, latitude string) (*Report, error)

// Report is the response that contains all Crimes information
type Report struct
    // String returns Report in readable format and receives the argument of Top number, if top = 0 => all data. Fe: top = 5 => 5 latest crimes at radious of coordinate
    String(top uint) string
```