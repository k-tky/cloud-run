package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Cloud Run running...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Cloud Run request recieve")
		str := os.Getenv("STRING")
		if str == "" {
			str = "STRING"
		}
		fmt.Printf("Hello Cloud Run! %s", str)
		w.Write([]byte(fmt.Sprintf("Hello Cloud Run! %s", str)))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
