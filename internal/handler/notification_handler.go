package handler

import (
	"encoding/json"
	"net/http"

	"notification-agent/internal/dto"
	"notification-agent/internal/service"
)

func CreateNotification(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request dto.CreateNotificationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	id, err := service.CreateNotification(
		request.UserID,
		request.Message,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"id":     id,
			"status": "CREATED",
		},
	)
}