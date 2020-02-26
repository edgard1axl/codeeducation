package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", IndexServer)
    http.ListenAndServe(":8000", nil)
}

func greetings(s string) string {
	return "<b>" + s + "</b>"
}

func IndexServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", greetings("Code.education Rocks !"))
}