package item

import (
	"log"
	"net/http"
	"time"

	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Item struct
type Item struct {
	ID         string    `json:"id"`
	PartNumber string    `json:"part_number"`
	Chemistry  string    `json:"chemistry"`
	CreatedAt  time.Time `json:"created_at"`
}

// Routes comment
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{itemID}", GetItem)
	router.Delete("/{itemID}", DeleteItem)
	router.Post("/", CreateItem)
	router.Get("/", GetAllItems)
	return router
}

// GetItem returns an item
func GetItem(w http.ResponseWriter, r *http.Request) {
	connStr := "postgresql://leader@localhost:26257/qc?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var item Item
	itemID := chi.URLParam(r, "itemID")
	err = db.QueryRow(`SELECT p.partnumber, p.chemistry, i.lotnumber FROM inventory AS i LEFT JOIN partnumbers AS p ON i.partnumberid = p.id WHERE p.partnumber = $1`, itemID).Scan(&item.ID, &item.Chemistry, &item.PartNumber)
	item.CreatedAt = time.Now()

	if err != nil {
		log.Fatal(err)
	}

	render.JSON(w, r, item)
}

// DeleteItem returns an item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response)
}

// CreateItem returns an item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response)
}

// GetAllItems returns an item
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := []Item{
		{
			ID:         "itemID",
			PartNumber: "Hello World",
			Chemistry:  "Hello World, from planet",
			CreatedAt:  time.Now(),
		},
	}
	render.JSON(w, r, items)
}
