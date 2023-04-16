package main

import (
	"github.com/6QHTSK/Bestdori-Proxy/config"
	"github.com/6QHTSK/Bestdori-Proxy/router"
	"log"
	"os"
)

func main() {
	r := router.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Bestdori-Proxy %s", config.Version)
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
