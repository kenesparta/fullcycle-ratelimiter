package handler

import (
	"encoding/json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	output := struct {
		Message string `json:"message"`
	}{
		Message: "Hello world!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
