package main

func main() {
	server, err := CreateServer()

	if err != nil {
		panic(err)
	}

	server.Version()
	server.ListRoute()
	server.Start()
}
