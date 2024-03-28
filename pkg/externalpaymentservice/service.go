package externalpaymentservice

import "log/slog"

func Pay(attributes string) error {
	slog.Info("Payment done",
		"attributes", attributes)

	return nil
}
