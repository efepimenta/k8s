package main

import (
	"fmt"
	"net/http"
)

func greeting(mes string) string {
    return "<b>" + mes + "</b>";
}

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

    out := `<!DOCTYPE html>
        <html lang="en">
        <head>
        <meta charset="UTF-8">
        <title>Go Server</title>
        </head>
        <body>`;
    put := `</body>
        </html>`;

		fmt.Fprintf(w, out + greeting("Code.education Rocks!") + put);
	})

	http.ListenAndServe(":8000", nil)
}
