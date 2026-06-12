package user

import (
	"fmt"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/alanis/anichat-backend/internal/features/auth"
	"github.com/alanis/anichat-backend/internal/generated/api"
	"github.com/alanis/anichat-backend/internal/utils"
	"github.com/oapi-codegen/runtime/types"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsersMe(
	ctx context.Context,
	request api.GetUsersMeRequestObject,
) (api.GetUsersMeResponseObject, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(string)
	log.Printf("Getting profile for user ID: %s", userID) // Логируем ID пользователя для отладки
	if !ok {
		return api.GetUsersMe401JSONResponse{
			Success: false,
			Message: utils.PtrString("user ID not found in context"),
		}, nil
	}

	usersMeData, err := h.service.GetUserProfile(ctx, userID)
	if err != nil {
		return api.GetUsersMe401JSONResponse{
			Success: false,
			Message: utils.PtrString("failed to get user profile"),
		}, nil
	}
	responseData := &api.UsersMeResponseData{
		Id:             fmt.Sprint(usersMeData.ID),
		Email:          types.Email(usersMeData.Email),
		ProfileCreated: usersMeData.ProfileCreated,
	}
	if usersMeData.ProfileCreated {
		responseData.Profile = &api.UsersMeProfile{
			FirstName: usersMeData.Profile.FirstName,
			LastName:  usersMeData.Profile.LastName,
			AvatarUrl: usersMeData.Profile.AvatarURL,
		}
	}
	return api.GetUsersMe200JSONResponse{
		Success: true,
		Data:    *responseData,
	}, nil
}

func (h *Handler) UpdateUsersMeProfile(
	ctx context.Context,
	request api.UpdateUsersMeProfileRequestObject,
) (api.UpdateUsersMeProfileResponseObject, error) {
	return nil, nil
}

func (h *Handler) GetUsersMe1(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(auth.UserIDKey).(string)
	log.Printf("Getting profile for user ID: %s", userID) // Логируем ID пользователя для отладки
	if !ok {
		http.Error(w, "user ID not found in context", http.StatusInternalServerError)
		return
	}

	response, err := h.service.GetUserProfile(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to get user profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpsertProfile(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := r.Context().Value(auth.UserIDKey).(string)

	if !ok {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	var req UpsertProfileRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	err := h.service.UpsertProfile(
		r.Context(),
		userID,
		req,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
