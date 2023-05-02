package ukcrimes

import "github.com/qazaqpyn/uk-crimes/internal/client"

func New() (*client.Client, error) {
	return client.NewClient(), nil
}
