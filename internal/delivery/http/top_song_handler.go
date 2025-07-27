package http

import (
	"PracticeCrud/internal/delivery/response"
	"PracticeCrud/internal/domain"
	"encoding/json"
	"net/http"
)

type TopSongHandler struct {
	Usecase domain.TopSongUsecase
}

func (h *TopSongHandler) Add(w http.ResponseWriter, r *http.Request) {
	var song domain.TopSong
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		response.WriterError(w, http.StatusBadRequest, "ошибка")
		return
	}
	if err := h.Usecase.Add(&song); err != nil {
		response.WriterError(w, http.StatusInternalServerError, "песня не добавлена")
		return
	}
	response.WriterSuccess(w, song, "песня добавлена в топ!")
}
func (h *TopSongHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	songs, err := h.Usecase.GetAll()
	if err != nil {
		response.WriterError(w, http.StatusNotFound, "песни не найдены")
		return
	}
	response.WriterSuccess(w, songs, "песни найдены")
}
