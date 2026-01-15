package main

import (
	"context"

	"github.com/RSODA/todo-go/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	err = a.RunHTTP()
	if err != nil {
		panic(err)
	}
}
