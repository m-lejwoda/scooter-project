package user

import (
	"net/http"

	"scooter-prj/internal/helper"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /user/login", h.Login)
	//mux.HandleFunc("POST /user/register", h.RegisterHTTP)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	u, err := helper.ReadJSON[User](w, r)
	if err != nil {
		return
	}
	if err := h.service.Login(r.Context(), u); err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "Error Server")
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"status": "Logged in"})
}
