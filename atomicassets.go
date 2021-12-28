package atomicassets

import (
	"encoding/json"
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
