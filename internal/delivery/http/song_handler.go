package http

import (
	"PracticeCrud/internal/delivery/response"
	"PracticeCrud/internal/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"strconv"
)

type SongHandler struct {
	Usecase domain.SongUsecase
}

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "songs_http_requests_total",
			Help: "Total number of HTTP requests to songs endpoints",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
}

func (h *SongHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues("GET", "/songs").Inc()
	songs, err := h.Usecase.GetAll()
	if err != nil {
		response.WriterError(w, http.StatusInternalServerError, "Не удалось получить песни")
		return
	}
	response.WriterSuccess(w, songs, "Список песен получен")
}

func (h *SongHandler) Get(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues("GET", "/songs/{id}").Inc()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.WriterError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}
	songs, err := h.Usecase.Get(id)
	if err != nil {
		response.WriterError(w, http.StatusNotFound, "Песня не найдена")
		return
	}
	response.WriterSuccess(w, songs, "Песня найдена")
}
func (h *SongHandler) Create(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues("POST", "/songs").Inc()
	var s domain.Song
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		response.WriterError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}
	if err := h.Usecase.Create(&s); err != nil {
		log.Println("Ошибка при создании:", err)
		response.WriterError(w, http.StatusInternalServerError, "Не удалось создать песню")
		return
	}
	response.WriterSuccess(w, s, "Песня успешно создана")
}
func (h *SongHandler) Update(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues("PUT", "/songs/{id}").Inc()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.WriterError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}
	var s domain.Song
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		response.WriterError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}
	s.ID = id
	if err = h.Usecase.Update(&s); err != nil {
		response.WriterError(w, http.StatusInternalServerError, "Не удалось обновить песню")
		return
	}
	response.WriterSuccess(w, s, "Песня обновлена")
}

func (h *SongHandler) Delete(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues("DELETE", "/songs/{id}").Inc()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.WriterError(w, http.StatusNotFound, "Некорректный ID")
		return
	}
	if err = h.Usecase.Delete(id); err != nil {
		response.WriterError(w, http.StatusInternalServerError, "error to delete")
		return
	}
	response.WriterSuccess(w, nil, "Песня удалена")
}
