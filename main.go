package main

import (
	"context"
	//"weak"
	"fmt"
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

type DynamicSearchResponse struct {
	Headers    []string
	Rows       []map[string]any
	NextOffset int
	Query      string
}

type SearchResponse struct {
	Headers     []string
	Rows        []map[string]any
	NextOffset  int
	Query       string
	CurrentSort string
	CurrentDir  string
} // Now

//go:embed web/templates/*.html
var templateFS embed.FS

func handlePlayersRequest(dbpool *pgxpool.Pool, tmpl *template.Template, w http.ResponseWriter, queryTerm string, offset int, templateName string, r *http.Request) {
	sortCol := r.URL.Query().Get("sort")
	sortDir := r.URL.Query().Get("dir")

	allowedSorts := map[string]string{
		"Name":  `"Name"`,
		"Fed":   `"Fed"`,
		"Tit":   `"Tit"`,
		"DEC25": `"DEC25"`,
		"B-day": `"B-day"`,
	}

	query := r.FormValue("search") // HTMX hx-post uses FormValue
	if query == "" {
		query = r.URL.Query().Get("search") // Load More and Sorting use URL Query
	}

	dbColumn, ok := allowedSorts[sortCol]
	if !ok {
		dbColumn = `"DEC25"`
	}
	if sortDir != "ASC" {
		sortDir = "DESC"
	}

	sql := fmt.Sprintf(`
        SELECT "B-day", "Name", "DEC25", "Fed", "Tit" 
        FROM fide_players 
        WHERE "Name" ILIKE $1 
        ORDER BY %s %s 
        LIMIT 20 OFFSET $2`, dbColumn, sortDir)

	// FIX: Ensure you use the variables returned here
	headers, rows, err := getDynamicPlayers(dbpool, sql, "%"+queryTerm+"%", offset)
	if err != nil {
		log.Println("Query Error:", err)
		http.Error(w, "Database error", 500)
		return
	}

	data := SearchResponse{
		Headers:     headers,
		Rows:        rows,
		NextOffset:  offset + 20,
		Query:       queryTerm,
		CurrentSort: sortCol,
		CurrentDir:  sortDir,
	}

	//'data' is defined and can be used
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

	// The Root Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Added 'r' at the end
		handlePlayersRequest(dbpool, tmpl, w, "", 0, "index.html", r)
	})

	// The Search Handler
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("search")
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		// Added 'r' at the end
		handlePlayersRequest(dbpool, tmpl, w, query, offset, "table.html", r)
	})
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
