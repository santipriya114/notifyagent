package repository

func GetNotificationStatus(id int64) (string, error) {

	var status string

	err := DB.QueryRow(
		"SELECT status FROM notifications WHERE id = $1",
		id,
	).Scan(&status)

	if err != nil {
		return "", err
	}

	return status, nil
}