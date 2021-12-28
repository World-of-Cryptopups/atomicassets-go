package atomicassets

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AtomicAssets struct {
	apiUrl string
	client *http.Client
}

func (a *AtomicAssets) _requester(u string, v interface{}) error {
	// request
	req, err := a.client.Get(u)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// Create a new AtomicAssets instance.
func New() *AtomicAssets {
	return &AtomicAssets{
		apiUrl: "https://wax.api.atomicassets.io/atomicassets/v1",
		client: &http.Client{},
	}
}

// Create a new AtomicAssets instance with custom api endpoint and http client.
func NewCustom(endpoint string, client *http.Client) (*AtomicAssets, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint provided is blank")
	}

	if client == nil {
		return nil, errors.New("custom http.Client is nil")
	}

	return &AtomicAssets{
		apiUrl: endpoint,
		client: client,
	}, nil
}
