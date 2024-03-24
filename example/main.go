package main

import (
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
}
