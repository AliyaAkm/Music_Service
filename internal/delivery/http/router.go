package http

import (
	"PracticeCrud/internal/domain"
	"github.com/gorilla/mux"
)

func RegisterSongRoutes(r *mux.Router, songUS domain.SongUsecase, topsongUS domain.TopSongUsecase) {
	handler := &SongHandler{Usecase: songUS}
	topHandler := &TopSongHandler{Usecase: topsongUS}
	r.HandleFunc("/songs", handler.Fetch).Methods("GET")
	r.HandleFunc("/songs/{id}", handler.Get).Methods("GET")
	r.HandleFunc("/songs", handler.Create).Methods("POST")
	r.HandleFunc("/songs/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/songs/{id}", handler.Delete).Methods("DELETE")
	// топ песен
	r.HandleFunc("/top_songs", topHandler.Add).Methods("POST")
	r.HandleFunc("/top_songs", topHandler.GetAll).Methods("GET")
}
