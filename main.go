package atomicassets

import "net/http"

// Create a new AtomicAssets instance.
func New() *AtomicAssets {
	return &AtomicAssets{
		API_URL: "https://wax.api.atomicassets.io/atomicassets/v1",
		client:  &http.Client{},
	}
}
