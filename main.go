package main

import (
	"context"
	//	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	//"strings"

	"embed"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Player struct {
	Bday  int
	Name  string
	DEC25 int
	Fed   string
	Tit   string
}

type SearchResponse struct {
	Players    []Player
	NextOffset int
	Query      string
}

//go:embed web/templates/*.html
var templateFS embed.FS

func getPlayers(dbpool *pgxpool.Pool, search string, offset int) []Player {
	query := `SELECT "B-day", "Name", "DEC25", "Fed", "Tit" 
              FROM fide_players 
              WHERE "Name" ILIKE $1 
              ORDER BY 3 DESC 
              LIMIT 20 OFFSET $2`

	rows, err := dbpool.Query(context.Background(), query, search, offset)
	if err != nil {
		log.Println("Query Error:", err)
		return nil // Return nil so the app doesn't crash
	}
	defer rows.Close()

	var searchResults []Player
	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.Bday, &p.Name, &p.DEC25, &p.Fed, &p.Tit); err != nil {
			log.Println("Scan Error:", err)
			continue
		}
		searchResults = append(searchResults, p)
	}
	return searchResults
}
func main() {
	connStr := "postgres://postgres:postgres@192.168.178.120:5432/test"
	//connStr := "postgres://postgres:postgres@localhost:5432/test"
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	//tmpl := template.Must(template.ParseGlob("web/templates/*.html"))
	tmpl := template.Must(template.ParseFS(templateFS, "web/templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Initialize the struct with an empty slice so len() doesn't crash
		initialData := SearchResponse{
			Players: []Player{},
		}

		err = tmpl.ExecuteTemplate(w, "index.html", initialData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("search")

		// Add this line to allow partial matches
		searchWildcard := "%" + query + "%"

		offsetStr := r.URL.Query().Get("offset")
		offset, _ := strconv.Atoi(offsetStr)
		if offsetStr == "" {
			offset = 0
		}

		// Use searchWildcard instead of query
		players := getPlayers(dbpool, searchWildcard, offset)

		data := SearchResponse{
			Players:    players,
			NextOffset: offset + 20,
			Query:      query, // Keep the clean query for the button URL
		}

		tmpl.ExecuteTemplate(w, "table.html", data)
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
