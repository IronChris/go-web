package main

import (
	"context"
	//"weak"
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
	Headers    []string         // Add this
	Rows       []map[string]any // Changed from []Player to []map[string]any
	NextOffset int
	Query      string
}

type DynamicSearchResponse struct {
	Headers    []string
	Rows       []map[string]any
	NextOffset int
	Query      string
}

//go:embed web/templates/*.html
var templateFS embed.FS

func handlePlayersRequest(dbpool *pgxpool.Pool, tmpl *template.Template, w http.ResponseWriter, queryTerm string, offset int, templateName string) {
	// The "Dynamic" SQL
	sql := `SELECT "B-day", "Name", "DEC25", "Fed", "Tit" 
            FROM fide_players 
            WHERE "Name" ILIKE $1 
            ORDER BY 3 DESC 
            LIMIT 20 OFFSET $2`

	// Add wildcards for the DB call
	headers, rows, err := getDynamicPlayers(dbpool, sql, "%"+queryTerm+"%", offset)
	if err != nil {
		log.Println("Query Error:", err)
		http.Error(w, "Database error", 500)
		return
	}

	data := SearchResponse{
		Headers:    headers,
		Rows:       rows,
		NextOffset: offset + 20,
		Query:      queryTerm,
	}

	err = tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Println("Template Error:", err)
	}
}

func getDynamicPlayers(dbpool *pgxpool.Pool, sql string, args ...any) ([]string, []map[string]any, error) {
	rows, err := dbpool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Extract column names from the result set
	fieldDescriptions := rows.FieldDescriptions()
	headers := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		headers[i] = string(fd.Name)
	}

	var results []map[string]any
	for rows.Next() {
		// values, err := rows.Values() returns []any for the current row
		values, err := rows.Values()
		if err != nil {
			return nil, nil, err
		}

		// Map each column name to its corresponding value
		rowMap := make(map[string]any)
		for i, colName := range headers {
			rowMap[colName] = values[i]
		}
		results = append(results, rowMap)
	}

	return headers, results, nil
}

func main() {

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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
		// Offset 0, empty search (%), renders index.html
		handlePlayersRequest(dbpool, tmpl, w, "", 0, "index.html")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// 1. Get the search term from POST or URL
		query := r.FormValue("search")

		// 2. Get the offset from the URL (for Load More)
		offsetStr := r.URL.Query().Get("offset")
		offset, _ := strconv.Atoi(offsetStr)

		// IMPORTANT: When HTMX sends a POST from the search box,
		// we want to return the WHOLE table (Headers + Rows).
		// When HTMX sends a GET from 'Load More', we only want the rows.

		handlePlayersRequest(dbpool, tmpl, w, query, offset, "table.html")
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
