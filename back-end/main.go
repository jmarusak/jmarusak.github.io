package main

import ( 
	"fmt"
	"log"
	"net/http"
)

var counter = 0

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprint(w, counter)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/counter", incrementCounter)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
