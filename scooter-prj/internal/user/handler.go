package user

import (
	"fmt"
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
	mux.HandleFunc("POST /user/register", h.Register)
	//mux.HandleFunc("POST /user/register", h.RegisterHTTP)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	u, err := helper.ReadJSON[UserLogin](w, r)
	if err != nil {
		return
	}
	userResp, err := h.service.Login(r.Context(), u)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "Error Server")
		return
	}
	fmt.Println(userResp)
	helper.WriteJSON(w, http.StatusOK, map[string]string{"status": "Logged in"})
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registration")
	u, err := helper.ReadJSON[UserRegister](w, r)
	if err != nil {
		return
	}
	user, err := h.service.Register(r.Context(), u)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "Error Server")
	}
	fmt.Println(user)
}
