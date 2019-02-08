package main

import (
	"fmt"
	"github.com/jackmrzhou/gc-ai/ab-app"
	"os"
)

func main() {
	var app ab_app.App
	if err := app.Initialize("/conf/app.conf"); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	app.Run()
}