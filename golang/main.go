package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// Item represents a simple data structure
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// In-memory data storage
var items []Item

// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// Get a single item by ID
func getItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, item := range items {
        if item.ID == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.NotFound(w, r)
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)
    item.ID = len(items) + 1
    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}

// Update an existing item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, item := range items {
        if item.ID == id {
            items = append(items[:index], items[index+1:]...)
            var updatedItem Item
            _ = json.NewDecoder(r.Body).Decode(&updatedItem)
            updatedItem.ID = id
            items = append(items, updatedItem)
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }
    http.NotFound(w, r)
}

// Delete an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, item := range items {
        if item.ID == id {
            items = append(items[:index], items[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(items)
}

func main() {
    // Initialize the router
    router := mux.NewRouter()

    // Seed with some data
    items = append(items, Item{ID: 1, Name: "Item 1"})
    items = append(items, Item{ID: 2, Name: "Item 2"})

    // Route handlers
    router.HandleFunc("/api/items", getItems).Methods("GET")
    router.HandleFunc("/api/items/{id}", getItem).Methods("GET")
    router.HandleFunc("/api/items", createItem).Methods("POST")
    router.HandleFunc("/api/items/{id}", updateItem).Methods("PUT")
    router.HandleFunc("/api/items/{id}", deleteItem).Methods("DELETE")

    // Start the server
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
