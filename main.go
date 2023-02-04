package main

func main() {
	server, err := CreateServer()

	if err != nil {
		panic(err)
	}

	server.Version()
	server.ListMiddleware()
	server.ListRoute()

	server.Start()
}
