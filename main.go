package main

import (

//	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"time"

	"github.com/wihdi/mnc/api"
)
func main() {
	// Inisialisasi router
	router := api.SetUpRouter()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server started at http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}


