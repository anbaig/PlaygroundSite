package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	thoughtsServer := http.FileServer(http.Dir("./thoughts"))
	staticServer := http.FileServer(http.Dir("./static"))
	http.Handle("/thoughts/", http.StripPrefix("/thoughts/", thoughtsServer))
	http.Handle("/", staticServer)
	
	port := 8080
	fmt.Printf("Starting server on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        log.Fatal(err)
    }
}