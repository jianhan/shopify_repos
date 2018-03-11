package main

import (
	"github.com/jianhan/shopify_repos/handlers"
)

func main() {
	handlers.Serve("127.0.0.1:8080")
}
