package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	service    *Service
	jwtManager *JwtManager
}

func NewHandler(service *Service, jwtManager *JwtManager) *Handler {
	return &Handler{service: service, jwtManager: jwtManager}
}

func (h *Handler) LoginByEmail(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	type response struct {
		Success  bool   `json:"success"`
		TicketId string `json:"ticketId,omitempty"`
		Message  string `json:"message,omitempty"`
	}
	ticketId, err := h.service.LoginByEmail(r.Context(), req.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "failed to send OTP", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{
			Success: false,
			Message: "failed to send OTP",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{
		Success:  true,
		TicketId: ticketId.String(),
		Message:  "OTP sent successfully",
	})
}

func (h *Handler) VerifyOtp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TicketId string `json:"ticketId"`
		Code     string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Received OTP verification request for %s with code %s", req.TicketId, req.Code) // Логируем входящие данные для отладки
	accessToken, refreshToken, err := h.service.VerifyLoginOtp(r.Context(), req.TicketId, req.Code)
	if err != nil {
		http.Error(w, "invalid OTP", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var res struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
	res.AccessToken = *accessToken
	res.RefreshToken = *refreshToken
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			claims, err := h.jwtManager.VerifyToken(tokenString)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			log.Printf("Authenticated user %s with token %s", claims.UserID, tokenString) // Логируем успешную аутентификацию для отладки
			ctx := context.WithValue(
				r.Context(),
				UserIDKey,
				claims.UserID,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
