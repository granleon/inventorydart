package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Item struct
type Item struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Manufacture time.Time     `bson:"manufacture" json:"manufacture"`
	Expire      time.Time     `bson:"expire" json:"expire"`
	LotNumber   string        `bson:"lotnumber" json:"lotnumber"`
}

// Incoming struct
type Incoming struct {
	ID          bson.ObjectId
	Manufacture string
	Expire      string
	LotNumber   string
}

// DB struct
type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func strToDate(str string) time.Time {
	month, _ := strconv.Atoi(str[2:4])
	day, _ := strconv.Atoi(str[4:])
	year, _ := strconv.Atoi(str[:2])
	return time.Date(year+2000, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (db *DB) getAllItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	w.WriteHeader(http.StatusOK)
	err := db.collection.Find(nil).All(&items)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(items)
		w.Write(response)
	}
}

func (db *DB) createItem(w http.ResponseWriter, r *http.Request) {
	var incoming Incoming
	itemBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(itemBody, &incoming)

	var item Item
	// create Hash ID for new item
	item.ID = bson.NewObjectId()
	item.Manufacture = strToDate(incoming.Manufacture)
	item.Expire = strToDate(incoming.Expire)

	err := db.collection.Insert(item)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(item)
		w.Write(response)
	}
}

func (db *DB) updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Item
	itemBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(itemBody, &item)
	// create a Hash ID
	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &item})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Update successfully"))
	}
}

func (db *DB) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Deleted sucessfully"))
	}
}

func (db *DB) getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post Item

	w.WriteHeader(http.StatusOK)
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&post)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(post)
		w.Write(response)
	}
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	collection := session.DB("qc").C("inventory")
	db := &DB{session: session, collection: collection}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[Connected] to MongoDB")

	defer session.Close()

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/item/{id}", db.getItem).Methods("GET")
	router.HandleFunc("/api/v1/item", db.createItem).Methods("POST")
	router.HandleFunc("/api/v1/item/{id}", db.updateItem).Methods("PATCH")
	router.HandleFunc("/api/v1/item/{id}", db.deleteItem).Methods("DELETE")
	router.HandleFunc("/api/v1/item", db.getAllItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", handlers.LoggingHandler(os.Stdout, handlers.CORS(originsOk, headersOk, methodsOk)(router))))
}
