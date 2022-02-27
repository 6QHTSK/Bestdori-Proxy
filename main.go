package main

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/router"
	"os"
)

func main() {
	r := router.InitRouter()
	port := os.Getenv("PORT")
	var err error
	if port == "" {
		err = r.Run("0.0.0.0:9000")
	} else {
		err = r.Run(":" + port)
	}
	if err != nil {
		return
	}
}
