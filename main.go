package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/metalstormbass/microservice-go/src/router"
)

const port = "8118"

func main() {
	r := router.Router()
	port_connect := fmt.Sprintf(":%s", port)
	log.Printf("HTTP Server listening on port %s", port)
	log.Print(http.ListenAndServe(port_connect, r))
}