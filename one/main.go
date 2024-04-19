package main

import (
	"context"
	"os"
)

func main() {
	component := hello("Ayush")
	component.Render(context.Background(), os.Stdout)
}
