package handler

import (
	"buf-validate/internal/application/service"
	"buf-validate/internal/infrastructure/repository"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var getTaskUsecase = service.GetTaskUsecase{
	TaskRepo: &repository.TaskRepositoryImpl{},
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := getTaskUsecase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
