package main

import (
	"backend-food/internal/api/router"
)

func main() {

	r := router.NewRouter()
	r.Engine.Run()

}
