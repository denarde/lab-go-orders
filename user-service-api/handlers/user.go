package handlers

import (
	"encoding/json"
	"net/http"
	"user-service-api/dto"
	"user-service-api/logger"
	"user-service-api/middlewares"
	"user-service-api/services"
	"user-service-api/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	requestID := middlewares.GetRequestID(r.Context())
	log := logger.WithRequestID(requestID)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error("Invalid request payload", err)
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := services.CreateUser(req)
	if err != nil {
		log.Error("Failed to create user", err)
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
