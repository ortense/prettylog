package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.WithValue(context.Background(), "key", false)

	type Person struct {
		Name string `json:"n"`
		Age  uint   `json:"a"`
	}

	logger.Error("This is a error")
	logger.Warn("Waaaaarning", "arg", Person{"me", 36}, "1", false)
	logger.Info("This is an info", "mult", "line")
	logger.Debug("Debug message")
	logger.ErrorContext(ctx, "error message")
}
