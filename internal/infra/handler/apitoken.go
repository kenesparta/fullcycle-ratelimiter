package handler

import (
	"encoding/json"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/usecase"
	"log"
	"net/http"
)

type APITokenHandler struct {
	repo entity.APITokenRepository
}

func NewAPITokenHandler(repo entity.APITokenRepository) *APITokenHandler {
	return &APITokenHandler{repo: repo}
}

func (at *APITokenHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
	input := dto.APITokenInput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("error decoding input data:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apiTokenUseCase := usecase.NewCreateAPITokenUseCase(at.repo)
	result, execErr := apiTokenUseCase.Execute(r.Context(), input)
	if execErr != nil {
		log.Println("error decoding input data:", execErr.Error())
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
