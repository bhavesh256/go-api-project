package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type listing struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Price       string    `json:"price"`
	City        string    `json:"city"`
	Created_at  time.Time `json:"created_at"`
}

type ListingHandler struct {
	db *sql.DB
	logger *slog.Logger
}

func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandler {
	return &ListingHandler{
		db: db,
		logger: logger,
	}
}

func (lh ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // fix zombie query problem
	rows, err := lh.db.QueryContext(
		ctx,
		`SELECT id, title, description, price, city, created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
	if err != nil {
		log.Printf("query: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	listings := []listing{}

	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Price, &l.City, &l.Created_at); err != nil {
			log.Printf("rows.Scan: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows,err: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(listings)
}

func (lh ListingHandler) DeleteListing(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	_, err := lh.db.ExecContext(
		ctx,
		`DELETE FROM listing WHERE id = $1`, id)

	if err != nil {
		log.Printf("delete: %v", err)
		lh.logger.Error("delete failed", "listing id", id, "error", err) // default : text handler
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("ok"))
}
