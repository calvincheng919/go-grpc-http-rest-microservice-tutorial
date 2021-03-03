package main

import (
	"fmt"
	"os"

	"github.com/calvincheng919/go-grpc-http-rest-microservice-tutorial/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.stderr, "%v\n", err)
		os.Exit(1)
	}
}
