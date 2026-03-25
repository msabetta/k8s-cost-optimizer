package main

import (
	"log"
	"net/http"
	"k8s-cost-optimizer/api"

)

func main() {
	router := api.SetupRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
