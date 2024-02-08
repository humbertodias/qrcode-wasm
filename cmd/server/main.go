package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "Port number")
	flag.Parse()
	fmt.Printf("Started on http://localhost:%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir("./assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
