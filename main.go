package main

import (
	"log"

	"github.com/code-katana/remitcompare/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
