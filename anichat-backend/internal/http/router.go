package http

import (
	"net/http"

	"github.com/alanis/anichat-backend/internal/features/auth"
	"github.com/alanis/anichat-backend/internal/features/user"
	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler *auth.Handler, userHandler *user.Handler) http.Handler {

	r := chi.NewRouter()

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/email", authHandler.LoginByEmail)
		r.Post("/email/otp", authHandler.VerifyOtp)
	})

	r.Group(func(r chi.Router) {
    r.Use(authHandler.AuthMiddleware())
    	r.Get("/api/users/me", userHandler.GetUsersMe)
    	r.Put("/api/users/me/profile", userHandler.UpsertProfile)
	})

	return r
}