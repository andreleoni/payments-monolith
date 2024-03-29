package externalpaymentservice

import (
	"log/slog"
	"payments/pkg/random"
)

func Pay(attributes string) (string, error) {
	slog.Info("Payment done on privider",
		"attributes", attributes,
		"provider", "fakeProviderName")

	return random.Hex(10), nil
}
