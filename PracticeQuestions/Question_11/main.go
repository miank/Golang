package main

import (
	"context"
	"fmt"
)

func greet(ctx context.Context) {

	name, ok := ctx.Value("name").(string)

	if !ok || name == "" {
		fmt.Println("Hello, guest!")
	} else {
		fmt.Printf("Hello %s! \n", name)
	}
}

func main() {

	ctx := context.WithValue(context.Background(), "name", "Alice")

	greet(ctx)

	ctxWithOutName := context.Background()
	greet(ctxWithOutName)

}
