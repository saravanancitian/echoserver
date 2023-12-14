package main

import (
	"echoserver/handlers"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("SERV_PORT")
	if !ok {
		fmt.Printf("Env not available\n")

		return
	}

	http.HandleFunc("/", handlers.GetRoot)
	http.HandleFunc("/hello", handlers.GetHello)
	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
