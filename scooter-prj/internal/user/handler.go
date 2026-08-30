package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"scooter-prj/internal/helper"
	"scooter-prj/internal/security"
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
	mux.HandleFunc("GET /user/test", h.UserTest)
	//mux.HandleFunc("POST /user/register", h.RegisterHTTP)
}

func (h *UserHandler) UserTest(w http.ResponseWriter, r *http.Request) {
	dict := map[string]string{
		"user":  "michal",
		"user1": "robert",
	}
	helper.WriteJSON(w, 200, dict)
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
	uID, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
	}
	userResp, err := h.service.RefreshToken(r.Context(), name, uID, username)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "Error Server")
	}
	helper.WriteJSON(w, http.StatusOK, &userResp)
}

func (h *UserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reset Password")
	u, _ := helper.ReadJSON[UserEmail](w, r)
	user, err := h.service.userRepo.GetByEmail(r.Context(), u.Email)
	if err != nil {
		fmt.Println("Error")
		if err == ErrUserNotFound {
			helper.WriteError(w, http.StatusBadRequest, "User not Found")
		}
	}

	hashedToken := security.GenerateSecureToken()
	resetTokenResponse, err := h.service.userRepo.SavePasswordResetToken(r.Context(), user.ID, hashedToken)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Something went wrong")
	}
	helper.WriteJSON(w, http.StatusOK, &resetTokenResponse)
}

//TODO Reset Password
//
//TODO Change Email with email send
//
//TODO read more features from prometheus
