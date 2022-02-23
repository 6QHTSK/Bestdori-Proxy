package main

import (
	"github.com/6QHTSK/ayachan-bestdoriAPI/router"
)

func main() {
	r := router.InitRouter()
	err := r.Run("0.0.0.0:21104")
	if err != nil {
		return
	}
}
