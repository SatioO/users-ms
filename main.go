package main

import "github.com/satioO/users/app"

func main() {
	server := app.Server{}
	server.Initialize()
	server.Run(":3000")
}
