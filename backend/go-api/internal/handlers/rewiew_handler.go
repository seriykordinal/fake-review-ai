package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seriykordinal/fake-review-ai/internal/dto"
	"github.com/seriykordinal/fake-review-ai/internal/services"
)

type ReviewHandler struct {
	service *services.ReviewService
}

func NewReviewHandler(s *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: s}
}

func (h *ReviewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Print(http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req dto.CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Print(http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.service.ProcessReview(r.Context(), req)
	if err != nil {
		log.Print(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
