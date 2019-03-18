package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
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
	w.Header().Set("content-type", "application/json")
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Unable to decode request body")
	}
	collection := client.Database("qc").Collection("inventory")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error connecting to DB")
	}
	result, err := collection.InsertOne(ctx, bson.M{
		"lotnumber":  item.LotNumber,
		"partnumber": item.PartNumber,
		"chem":       item.Chem,
		"chemabbr":   item.ChemAbbr,
		"expire":     item.Expire,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error retrieving record")
	}
	json.NewEncoder(w).Encode(result)
}

// GetItem creates an item
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	fmt.Println(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Unable to get id from params")
	}
	var item Item
	collection := client.Database("qc").Collection("inventory")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if cancel != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error connecting to DB")
	}
	err = collection.FindOne(ctx, Item{ID: id}).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error retrieving record")
	}
	json.NewEncoder(w).Encode(item)
}

// GetItems creates an item
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var items []Item
	collection := client.Database("qc").Collection("inventory")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if cancel != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error connecting to DB")
	}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error retrieving record")
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item Item
		cursor.Decode(&item)
		items = append(items, item)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
	}
	json.NewEncoder(w).Encode(items)
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
	router.HandleFunc("/api/v1/item/{id}", GetItem).Methods("GET")
	router.HandleFunc("/api/v1/item", CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{id}", UpdateItem).Methods("PATCH")
	router.HandleFunc("/api/v1/item/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v1/item", GetItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", router))
}
