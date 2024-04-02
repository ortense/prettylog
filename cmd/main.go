package main

import (
	"bufio"
	"os"

	"github.com/ortense/prettylog"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		prettylog.Print(line)
	}

	if err := scanner.Err(); err != nil {
		return
	}
}
