package atomicassets

// Create a new AtomicAssets instance.
func New() *AtomicAssets {
	return &AtomicAssets{
		Url: "https://wax.api.atomicassets.io/atomicassets/v1",
	}
}
