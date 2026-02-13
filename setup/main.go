package main

import (
	"golang_learning_blog/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run()
}
