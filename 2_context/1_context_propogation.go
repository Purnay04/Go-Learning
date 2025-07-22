package main

import (
	"context"
	"fmt"
)

func printUser(ctx context.Context) {
	// read "user" value from context
	user := ctx.Value("user")
	fmt.Println("User from context:", user)
}

func main() {
	// Create a background context
	ctx := context.Background()

	// Add a value to the context — propagate "user" key
	ctxWithUser := context.WithValue(ctx, "user", "john doe")

	// Pass the context to a function (could be a goroutine)
	printUser(ctxWithUser)
}
