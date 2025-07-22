package handler

import (
	"encoding/json"
	"net/http"

	"github.com/w1ntermute14/labhub-proxy/internal/models"
	"github.com/w1ntermute14/labhub-proxy/internal/services"
)

type Handler struct {
	service *services.ProxyService
}

func NewHandler(service *services.ProxyService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleHba1c(w http.ResponseWriter, r *http.Request) {
	var req models.Hba1cRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	prediction, err := h.service.PredictHba1c(req)
	if err != nil {
		http.Error(w, "Ошибка вычисления предсказания", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prediction)
}

func (h *Handler) HandleLdll(w http.ResponseWriter, r *http.Request) {
	var req models.LdllRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	prediction, err := h.service.PredictLdll(req)
	if err != nil {
		http.Error(w, "Ошибка вычисления предсказания", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prediction)
}

func (h *Handler) HandleFerr(w http.ResponseWriter, r *http.Request) {
	var req models.FerrRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	prediction, err := h.service.PredictFerr(req)
	if err != nil {
		http.Error(w, "Ошибка вычисления предсказания", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prediction)
}

func (h *Handler) HandleLdl(w http.ResponseWriter, r *http.Request) {
	var req models.LdlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	prediction, err := h.service.PredictLdl(req)
	if err != nil {
		http.Error(w, "Ошибка вычисления предсказания", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prediction)
}

func (h *Handler) HandleTg(w http.ResponseWriter, r *http.Request) {
	var req models.TgRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	prediction, err := h.service.PredictTg(req)
	if err != nil {
		http.Error(w, "Ошибка вычисления предсказания", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prediction)
}

func (h *Handler) HandleHdl(w http.ResponseWriter, r *http.Request) {
	var req models.HdlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	prediction, err := h.service.PredictHdl(req)
	if err != nil {
		http.Error(w, "Ошибка вычисления предсказания", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(prediction)
}
