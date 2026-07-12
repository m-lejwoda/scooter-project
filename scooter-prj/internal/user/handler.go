package user

import "net/http"

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /user/login", h.service.Login)
	//mux.HandleFunc("POST /user/register", h.RegisterHTTP)
}
