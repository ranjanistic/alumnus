package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Root")
	})
	
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "profile")
	})
	
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "dash")
    })

	fmt.Printf("Starting server at port 3000\n")
    if err := http.ListenAndServe("localhost:3000", nil); err != nil {
        log.Fatal(err)
	}
}