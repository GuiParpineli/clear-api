package main

import (
	"clear-api/internal/handler"
)

func main() {
	r := handler.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
