package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageData struct {
	Title string
	Items []string
}

func main() {
	// Initialisiere eine leere Liste für Einträge
	entries := []string{}

	// Handler für die Hauptseite
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := PageData{
			Title: "HTMX Go Beispiel",
			Items: entries,
		}
		tmpl.Execute(w, data)
	})

	// Handler für neue Einträge (HTMX Endpoint)
	http.HandleFunc("/add-entry", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Formularwert auslesen
		r.ParseForm()
		newEntry := strings.TrimSpace(r.FormValue("entry"))

		if newEntry != "" {
			// Neuen Eintrag zur Liste hinzufügen
			entries = append(entries, newEntry)

			// HTML-Template für den neuen Eintrag rendern
			tmpl := template.Must(template.New("entry").Parse(`
                <div class="entry-item" id="entry-{{len .Items}}">
                    <span>{{.NewEntry}}</span>
                    <button class="delete-btn" 
                            hx-delete="/delete-entry/{{len .Items}}"
                            hx-target="#entry-{{len .Items}}"
                            hx-swap="outerHTML">
                        ✕
                    </button>
                </div>
            `))

			data := struct {
				NewEntry string
				Items    []string
			}{
				NewEntry: newEntry,
				Items:    entries,
			}

			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, data)
		} else {
			fmt.Fprint(w, "<div class='error'>Bitte geben Sie einen Text ein</div>")
		}
	})

	// Handler zum Löschen von Einträgen
	http.HandleFunc("/delete-entry/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// ID aus URL extrahieren (vereinfacht für dieses Beispiel)
		// In einer echten App würde man hier die ID parsen und aus der Liste entfernen

		// Leere Antwort = Element wird entfernt (durch HTMX outerHTML swap)
		fmt.Fprint(w, "")
	})

	// Statische Dateien
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server läuft auf http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
