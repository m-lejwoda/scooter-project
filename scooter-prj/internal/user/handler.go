package user

import (
	"errors"
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
	mux.HandleFunc("POST /user/refresh_token", h.RefreshToken)
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

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Refresh token")
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Println("No cookie")
		}
	}
	refreshToken := cookie.Value

	claims, _ := VerifyToken(refreshToken)
	userID, _ := claims["user_id"].(string)
	username, _ := claims["username"].(string)

	name := username + userID
	h.service.RefreshToken(r.Context(), name)
	// TODO FINISH THIS
}
