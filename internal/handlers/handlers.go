func ListProducts(db *pgxpool.Pool, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, _ := db.Query(r.Context(), "SELECT \"B-day\", \"Name\", \"DEC25\" from  fide_players limit 5")
		var products []models.Product

		for rows.Next() {
			var p models.Product
			rows.Scan(&p.B-day, &p.Name, &p.DEC25)
			products = append(products, p)
		}

		// Render the template with the data
		tmpl.ExecuteTemplate(w, "table.html", products)
	}
}
