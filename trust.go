package main

import (
	"example/anon/trust/deal"
	"example/anon/trust/server"
	"example/anon/trust/test"
	"fmt"

	"github.com/gorilla/sessions"
)

var archives deal.Archives

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key          = []byte("super-secret-key")
	sessionStore = sessions.NewCookieStore(key)
)

func main() {
	fmt.Println("PROGRAM START")

	archives = deal.InitArchives()

	test.InitData(&archives)

	//test.All(&archives)

	//jeison.TestJson()

	server.Server(&archives)

	//goserv()

	println("\nPROGRAM END")
}
