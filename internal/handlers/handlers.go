package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"prod/internal/models"
	"prod/internal/service"
)

type Handler struct {
	Store   *service.DataStore
	BaseURL string
}

func NewHandler(store *service.DataStore, baseURL string) *Handler {
	return &Handler{
		Store:   store,
		BaseURL: baseURL,
	}
}

func ResponseJson(w http.ResponseWriter, statuscode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func ResponseError(w http.ResponseWriter, statuscode int, message string) {
	ResponseJson(w, statuscode, models.ErrorRes{Error: message})
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	req := models.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Store.Create(r.Context(), req.Task, req.Description)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	tasks, err := h.Store.GetAll(r.Context())
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, tasks)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	id := models.GetByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	task, err := h.Store.GetByID(r.Context(), id.ID)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, task)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	req := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Store.Update(r.Context(), req.ID, req.Task, req.Description, req.Do)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, map[string]string{"message": "Task updated successfully"})
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	id := models.GetByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Store.Delete(r.Context(), id.ID)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, map[string]string{"message": "Task deleted successfully"})
}

func Test(w http.ResponseWriter, r *http.Request) {
	ResponseJson(w, http.StatusOK, map[string]string{"Method": r.Method,
		"Path":  r.URL.Path,
		"Query": r.URL.RawQuery,
		"HI?":   "Hello World"})
}
