	package handler

	import (
		"encoding/json"
		"net/http"

		"notification-agent/internal/repository"
	)

	func GetNotificationsHandler(
		w http.ResponseWriter,
		r *http.Request,
	) {

		id := r.URL.Query().Get("id")
		status := r.URL.Query().Get("status")

		println("ID =", id)
		println("STATUS =", status)

		// GET /notifications?status=PENDING
		if status != "" {

			notifications, err := repository.GetNotificationsByStatus(status)

			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			json.NewEncoder(w).Encode(notifications)

			return
		}

		// GET /notifications?id=1
		if id != "" {

			notification, err := repository.GetNotificationByID(id)

			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusNotFound,
				)
				return
			}

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			json.NewEncoder(w).Encode(notification)

			return
		}

		// GET /notifications
		notifications, err := repository.GetNotifications()

		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		json.NewEncoder(w).Encode(notifications)
	}