package handler

import (
	"encoding/json"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"net/http"
)

type IPHandler struct {
	ipRepo entity.IPRepository
}

func NewIPHandler(ipRepo entity.IPRepository) *IPHandler {
	return &IPHandler{ipRepo: ipRepo}
}
func (h *IPHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Message string `json:"message"`
	}{
		Message: "hello world!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
