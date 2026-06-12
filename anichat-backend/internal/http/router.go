package http

import (
	"context"
	"net/http"

	"github.com/alanis/anichat-backend/internal/features/auth"
	"github.com/alanis/anichat-backend/internal/features/user"
	api "github.com/alanis/anichat-backend/internal/generated/api"
	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler *auth.Handler, userHandler *user.Handler) http.Handler {

	r := chi.NewRouter()

	// r.Route("/api/auth", func(r chi.Router) {
	// 	r.Post("/email", authHandler.LoginByEmail)
	// 	r.Post("/email/otp", authHandler.VerifyOtp)
	// })

	// r.Group(func(r chi.Router) {
	// r.Use(authHandler.AuthMiddleware())
	// 	r.Get("/api/users/me", userHandler.GetUsersMe)
	// 	r.Put("/api/users/me/profile", userHandler.UpsertProfile)
	// })

	handler := NewHandler(authHandler, userHandler)

	strict := api.NewStrictHandler(
		handler,
		[]api.StrictMiddlewareFunc{
			authHandler.JWTStrictMiddleware,
		},
	)

	api.HandlerFromMux(strict, r)

	return r
}

type Handler struct {
	authHandler *auth.Handler
	userHandler *user.Handler
}

func NewHandler(authHandler *auth.Handler, userHandler *user.Handler) *Handler {
	return &Handler{
		authHandler: authHandler,
		userHandler: userHandler,
	}
}

func (h *Handler) AuthByEmail(
	ctx context.Context,
	request api.AuthByEmailRequestObject,
) (api.AuthByEmailResponseObject, error) {
	return h.authHandler.AuthByEmail(ctx, request)
}

func (h *Handler) VerifyEmailOtp(
	ctx context.Context,
	request api.VerifyEmailOtpRequestObject,
) (api.VerifyEmailOtpResponseObject, error) {
	return h.authHandler.VerifyEmailOtp(ctx, request)
}

func (h *Handler) GetUsersMe(
	ctx context.Context, 
	request api.GetUsersMeRequestObject,
) (api.GetUsersMeResponseObject, error) {
	return h.userHandler.GetUsersMe(ctx, request)
}

func (h *Handler) UpdateUsersMeProfile(
	ctx context.Context, 
	request api.UpdateUsersMeProfileRequestObject,
) (api.UpdateUsersMeProfileResponseObject, error) {
	return h.userHandler.UpdateUsersMeProfile(ctx, request)
}

