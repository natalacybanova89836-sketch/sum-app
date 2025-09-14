package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "param 'a' missing or not integer", http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "param 'b' missing or not integer", http.StatusBadRequest)
		return
	}

	res := Add(a, b)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]int{"sum": res})
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	addr := ":8080"
	fmt.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
