package main

import "key-value-storage/app/server"

func main() {
	app := server.NewApp()
	app.Start()
}
