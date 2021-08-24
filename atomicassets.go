package atomicassets

import (
	"encoding/json"
	"net/http"
)

type AtomicAssets struct {
	API_URL string
	client  *http.Client
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
