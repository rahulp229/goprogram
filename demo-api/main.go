package main

import "demo-api/routers"

func main() {
	routes := routers.Routes()

	routes.Run(":8000")
}
