package main

//	@title			Open Courseware Application
//	@version		1.0.1
//	@description	This is Open Couseware backend
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
