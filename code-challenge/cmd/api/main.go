package main

import "github.com/iofabela/technical-challenge-meli/cmd/api/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
