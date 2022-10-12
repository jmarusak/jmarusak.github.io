package main

import ( 
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
	"cloud.google.com/go/datastore"
)

type Counter struct {
	Visits int
	LastVisit time.Time
}

func getCounterFromDatastore() Counter {
	ctx := context.Background()

	projectID := "martinview4"

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal("Datastore - Failed to create client %v", err)
	}
	defer client.Close()

	kind := "Counter"
	name := "counter"
	counterKey := datastore.NameKey(kind, name, nil)

	var counter Counter
	if err := client.Get(ctx, counterKey, &counter); err !=nil && err != datastore.ErrNoSuchEntity {	
		log.Fatalf("Datastore - Failed to get Counter: %v", err)
	}
	log.Println(counter)

	currentCounter := counter

	counter.Visits += 1
	counter.LastVisit = time.Now()

	if _, err := client.Put(ctx, counterKey, &counter); err != nil {
         log.Fatalf("Datastore - Failed to save Counter: %v", err)
	}

	return currentCounter
}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Counter REST API")
}

func serveCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "null")

	counter := getCounterFromDatastore()
	err := json.NewEncoder(w).Encode(counter)
    if err != nil {
		log.Fatal("Error building respose %v", err)
	}
}

func main() {
	http.HandleFunc("/", serveHomePage)
	http.HandleFunc("/counter", serveCounter)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
