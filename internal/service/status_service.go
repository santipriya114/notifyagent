package service

import "notification-agent/internal/repository"

func GetNotificationStatus(id int64) (string, error) {
	return repository.GetNotificationStatus(id)
}