package main

type client interface {
	Send()
	Close()
}

func main() {
	SendNRequestsWithMClients(100, 100000)
}
