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

var globalFederations []string

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
	CurrentFed  string   // New
	Federations []string // New: List for the dropdown
	TotalCount  int
}

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

	/*	query := r.FormValue("search") // HTMX hx-post uses FormValue
		if query == "" {
			query = r.URL.Query().Get("search") // Load More and Sorting use URL Query
		}
	*/
	dbColumn, ok := allowedSorts[sortCol]
	if !ok {
		dbColumn = `"DEC25"`
	}
	if sortDir != "ASC" {
		sortDir = "DESC"
	}

	fed := r.FormValue("fed")
	if fed == "" {
		fed = r.URL.Query().Get("fed")
	}

	// Build the base query
	sql := fmt.Sprintf(`
    SELECT "B-day", "Name", "DEC25", "Fed", "Tit", COUNT(*) OVER() as total_count
    FROM fide_players 
    WHERE "Name" ILIKE $1 
    AND ($3 = '' OR "Fed" = $3) 
    ORDER BY %s %s 
    LIMIT 20 OFFSET $2`, dbColumn, sortDir)

	// 2. Your function call MUST match the placeholders ($1, $2, $3):
	// $1 = name search, $2 = offset, $3 = fed
	headers, rows, totalCount, err := getDynamicPlayers(dbpool, sql, "%"+queryTerm+"%", offset, fed)
	// Inside handlePlayersRequest
	if err != nil {
		log.Println("Query Error:", err)
		http.Error(w, "Database error", 500)
		return
	}

	data := SearchResponse{
		Headers:     headers,
		Rows:        rows,
		TotalCount:  totalCount,
		NextOffset:  offset + 20,
		Query:       queryTerm,
		CurrentSort: sortCol,
		CurrentDir:  sortDir,
		CurrentFed:  fed,
		Federations: globalFederations, // Pass the global list here
	}
	//'data' is defined and can be used
	err = tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Println("Template Error:", err)
	}
}

func getFederations(db *pgxpool.Pool) []string {
	ctx := context.Background()
	rows, err := db.Query(ctx, `SELECT DISTINCT "Fed" FROM fide_players WHERE "Fed" IS NOT NULL ORDER BY "Fed" ASC`)
	if err != nil {
		return []string{}
	}
	defer rows.Close()

	var feds []string
	for rows.Next() {
		var f string
		rows.Scan(&f)
		feds = append(feds, f)
	}
	return feds
}

func getDynamicPlayers(dbpool *pgxpool.Pool, sql string, args ...interface{}) ([]string, []map[string]any, int, error) {
	rows, err := dbpool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, nil, 0, err
	}
	defer rows.Close()

	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, fd := range fieldDescriptions {
		columns = append(columns, string(fd.Name))
	}

	var resultRows []map[string]any
	totalCount := 0

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("scan erroe", err)
			return nil, nil, 0, err
		}

		rowMap := make(map[string]any)
		for i, colName := range columns {
			if colName == "total_count" {
				// Window function returns int64 in pgx
				if v, ok := values[i].(int64); ok {
					totalCount = int(v)
				}
				continue
			}
			rowMap[colName] = values[i]
		}
		resultRows = append(resultRows, rowMap)
	}

	// FIX: Properly build the display headers by excluding total_count
	var displayHeaders []string
	for _, h := range columns {
		if h != "total_count" {
			displayHeaders = append(displayHeaders, h) // Added 'h' here!
		}
	}

	// ... existing loop ...
	fmt.Printf("DEBUG: Headers: %v\n", displayHeaders)
	if len(resultRows) > 0 {
		fmt.Printf("DEBUG: Sample Row: %v\n", resultRows[0])
	}

	return displayHeaders, resultRows, totalCount, nil
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

	// Fetch this once when the server starts
	globalFederations = getFederations(dbpool)

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
		// r.FormValue captures BOTH the typing (POST) and the Sort/Load More (GET)
		searchTerm := r.FormValue("search")

		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		handlePlayersRequest(dbpool, tmpl, w, searchTerm, offset, "table.html", r)
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
