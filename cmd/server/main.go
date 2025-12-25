package main

import (
	"context"
	//	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Players struct {
	Bday  int
	Name  string
	DEC25 int
}

func main() {
	connStr := "postgres://postgres:postgres@192.168.178.120:5432/test"
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	tmpl := template.Must(template.ParseGlob("web/templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := dbpool.Query(context.Background(), "SELECT \"B-day\", \"Name\", \"DEC25\" from fide_players limit 10")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var players []Players
		for rows.Next() {
			var p Players
			if err := rows.Scan(&p.Bday, &p.Name, &p.DEC25); err != nil {
				log.Println(err)
				continue
			}
			players = append(players, p)
		}
		tmpl.ExecuteTemplate(w, "index.html", players)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")
		if !strings.Contains(search, "%") {
			search = search + "%"
		}

		// FIX 1: You need to perform the query and SCAN the results inside THIS function
		rows, err := dbpool.Query(context.Background(),
			"SELECT \"B-day\", \"Name\", \"DEC25\" from fide_players WHERE \"Name\" ILIKE $1 LIMIT 20",
			search)
		if err != nil {
			log.Println(err)
			return
		}
		defer rows.Close()

		var searchResults []Players // Local slice for this handler
		for rows.Next() {
			var p Players
			if err := rows.Scan(&p.Bday, &p.Name, &p.DEC25); err != nil {
				log.Println(err)
				continue
			}
			searchResults = append(searchResults, p)
		}

		// FIX 2: Use searchResults here
		// Note: We use "table-rows" to return JUST the rows for HTMX
		tmpl.ExecuteTemplate(w, "table-rows", searchResults)
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
