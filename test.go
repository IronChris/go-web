package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func readCheatSheet(w http.ResponseWriter, req *http.Request) {

	path := "/home/chris/projects/CHEAT_SHEET.md"
	dat, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	// Output file content correctly
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(dat)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/cs", readCheatSheet)

	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8092", nil)
}

