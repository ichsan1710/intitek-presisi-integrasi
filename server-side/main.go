package main

import (
	"log"
	"net/http"
	"server-side/config"
	"server-side/controllers/productcontroller"

	"github.com/rs/cors"
)

func main() {
	config.ConnectDB()

	// Products

	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/update", productcontroller.Update)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	// CORS Configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	handler := c.Handler(http.DefaultServeMux)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", handler)
}
