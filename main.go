package main

import (
	"cloud-compute/bootstrap"
)

func main() {
	app := bootstrap.Boot()

	app.Start()
}
