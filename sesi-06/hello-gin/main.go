package main

import Routers "hello/routers"

func main() {
	var PORT = ":8080"

	Routers.StartServer().Run(PORT)
}
