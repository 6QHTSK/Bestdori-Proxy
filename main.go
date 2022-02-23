package main

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/router"
	"os"
)

func main() {
	r := router.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "21104"
	}
	err := r.Run("0.0.0.0:" + port)
	if err != nil {
		return
	}
}
