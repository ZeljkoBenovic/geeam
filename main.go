package main

import (
	"log"

	"github.com/ZeljkoBenovic/geeam/core/engine"
)

func main() {
	// bootstrap main application
	mainApp, err := engine.BootstrapCoreApp()
	if err != nil {
		log.Fatalln("Could not bootstrap main app:", err.Error())
	}

	// run the app
	err = mainApp.Run()
	if err != nil {
		log.Fatalln("Could not run:", err.Error())
	}
}
