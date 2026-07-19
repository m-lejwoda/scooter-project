package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func GetUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, ok := GetTokenFromRequest(r)
		if !ok {
			http.Error(w, "Incorrect Header", http.StatusUnauthorized)
			return
		}

		claims, err := VerifyAccessToken(token)
		if err != nil {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Wrong token structure", http.StatusUnauthorized)
			return
		}

		userID := int(userIDFloat)
		fmt.Println("User Verified. ID:", userID)
		type contextKey string
		const UserIDKey contextKey = "UserID"
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
