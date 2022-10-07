package main

import Routers "restapi/routers"

func main() {
	// declare port
	var PORT = ":3000"

	// call router func
	Routers.StartServer().Run(PORT)
}
