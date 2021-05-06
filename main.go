package main

import (
	"logger-viewer/routers"
)

func main() {
	router := routers.SetupRouter()

	err := router.Run(":9000")
	if err != nil {
		panic(err)
	}

}
