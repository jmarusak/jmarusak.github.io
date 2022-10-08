package main

import ( 
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
)

type Counter struct {
	Visits int
	LastVisit time.Time
}

var counter = Counter {
	Visits: 0,
	LastVisit: time.Now(),
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "null")

	err := json.NewEncoder(w).Encode(counter)
    if err != nil {
		log.Fatal("Error building respose %v", err)
	}

	// incrememt counter:w
	counter.Visits += 1
	counter.LastVisit = time.Now()
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/counter", incrementCounter)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
