package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"

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
	LotNumber  string             `json:"lotnumber" bson:"lotnumber,omitempty"`
	PartNumber string             `json:"partnumber" bson:"partnumber,omitempty"`
	Chem       string             `json:"chem" bson:"chem,omitempty"`
	ChemAbbr   string             `json:"chemabbr" bson:"chemabbr,omitempty"`
	Expire     string             `json:"expire" bson:"expire,omitempty"`
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

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
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var item Item
	collection := client.Database("qc").Collection("inventory")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctx, Item{ID: id}).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
	}
	json.NewEncoder(w).Encode(item)
}

// GetItems creates an item
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var items []Item
	collection := client.Database("qc").Collection("inventory")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	fmt.Println("Connected to MongoDB...")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/item/{id}", GetItem).Methods("GET")
	router.HandleFunc("/api/v1/item", CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{id}", UpdateItem).Methods("PATCH")
	router.HandleFunc("/api/v1/item/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/v1/item", GetItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", handlers.LoggingHandler(os.Stdout, handlers.CORS(originsOk, headersOk, methodsOk)(router))))
}
