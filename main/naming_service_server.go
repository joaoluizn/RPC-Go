package main

func name_service_main() {
	namingServiceServer := naming.NewNamingServiceServer("4769")
	namingServiceServer.Run()
}
