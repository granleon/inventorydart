package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Item struct
type Item struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	LotNumber  string             `json:"lotnumber" bson:"lotnumber"`
	PartNumber string             `json:"partnumber" bson:"partnumber"`
	Chem       string             `json:"chem" bson:"chem"`
	ChemAbbr   string             `json:"chemabbr" bson:"chemabbr"`
	Expire     string             `json:"expire" bson:"expire"`
}

// CreateItem creates an item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateItem placeholder")
}

// GetItem creates an item
func GetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetItem placeholder")
}

// GetItems creates an item
func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetItems placeholder")
}

// UpdateItem creates an item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateItem placeholder")
}

// DeleteItem creates an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteItem placeholder")
}

func main() {
	fmt.Println("Starting the application....on port localhost:9000")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/item/{itemID}", GetItem).Methods("GET")
	router.HandleFunc("/api/v1/item", CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{itemID}", UpdateItem).Methods("PATCH")
	router.HandleFunc("/api/v1/item/{itemID}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v1/item", GetItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", router))
}
