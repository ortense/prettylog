package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	type Nested struct {
		Foo string `json:"foo"`
		Bar uint   `json:"bar"`
	}

	logger.Info("Heeey!")
	logger.Error("Booooom!", "nested", Nested{"baz", 42}, "bool", false)
	logger.Warn("Keep", "in", "alert")
	logger.Info("This is an info", "mult", "line")
	logger.Debug("Debug message")
	fmt.Println("This is a simple text message")
	fmt.Println("{\"time\":\"2024-03-28T12:30:25.627417-03:00\",\"level\":\"unknown level\",\"msg\":\"test\",\"prop-key\":\"prop-value\"}")
	fmt.Println("{\"level\":\"warn\",\"msg\":\"this log has no time property\"}")
	fmt.Println("{\"time\":\"2024-03-28T12:30:25.627417-03:00\",\"text\":\"this log has no level and msg properties\"}")
}
