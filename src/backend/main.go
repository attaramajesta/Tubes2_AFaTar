package main

import (
	"net/http"
	"github.com/rs/cors"
	"example.com/myproject/ids"
	"example.com/myproject/bfs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ids", ids.IDSHandler)
	mux.HandleFunc("/bfs", bfs.BFSHandler)

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(mux)

	http.ListenAndServe(":8080", handler)
}