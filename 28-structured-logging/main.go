package main

import (
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic("Could not open file")
	}

	logger := slog.New(
		slog.NewJSONHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)

	slog.SetDefault(logger)

	slog.Info("this is INFO log")
	slog.Warn("this is WARN log")
	slog.Error("this is ERROR log")
	slog.Debug("this is DEBUG log")

	slog.Info(
		"this is complex log",
		slog.String("accepted_values", "key/value pairs with specific types for marshalling"),
		slog.Int("an int: ", 30),
		slog.Group("grouped_info", slog.String("you_can", "do this too")),
	)
}
