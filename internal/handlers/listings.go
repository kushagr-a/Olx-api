package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

// dependency injection using wrapper function.
// clouser factory
func Listlisting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(
			`SELECT id,title,price,description,city, created_at FROM listings 
			ORDER BY created_at DESC 
			LIMIT 20`,
		)
		if err != nil {
			log.Printf("Error in Query :%v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		listings := []listing{}
		for rows.Next() {
			var l listing
			if err := rows.Scan(
				&l.ID,
				&l.Title,
				&l.Price,
				&l.Description,
				&l.City,
				&l.CreatedAt,
			); err != nil {
				log.Printf("Error scanning row: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			listings = append(listings, l)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Error in scanning the rows: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(listings); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

	}
}
