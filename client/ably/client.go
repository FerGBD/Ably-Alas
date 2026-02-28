package ably

import (
	"os"

	"github.com/ably/ably-go/ably"
)

func NewClient() (*ably.Realtime, error) {
	apiKey := os.Getenv("ABLY_API_KEY")
	return ably.NewRealtime(
		ably.WithKey(apiKey),
		ably.WithClientID("backend-go-pro"),
	)
}
