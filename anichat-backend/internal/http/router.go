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

