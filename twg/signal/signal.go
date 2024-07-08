package signal

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Name:       "John Doe",
		Age:        25,
		Occupation: "gardener",
	}
	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
