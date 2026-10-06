package repository

func CreateNotification(userID string, message string) (int64, error) {

	query := `
	INSERT INTO notifications(
		user_id,
		channel,
		message,
		status
	)
	VALUES(
		$1,
		'EMAIL',
		$2,
		'PENDING'
	)
	RETURNING id
	`

	var id int64

	err := DB.QueryRow(
		query,
		userID,
		message,
	).Scan(&id)

	return id, err
}

func SaveNotification(
	channel string,
	message string,
	status string,
) error {

	query := `
	INSERT INTO notification_history
	(channel, message, status)
	VALUES ($1, $2, $3)
	`

	_, err := DB.Exec(
		query,
		channel,
		message,
		status,
	)

	return err
}

func GetNotifications() ([]map[string]interface{}, error) {

	rows, err := DB.Query(`
		SELECT id, user_id, channel, message, status
		FROM notifications
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notifications []map[string]interface{}

	for rows.Next() {

		var id int64
		var userID string
		var channel string
		var message string
		var status string

		err := rows.Scan(
			&id,
			&userID,
			&channel,
			&message,
			&status,
		)

		if err != nil {
			return nil, err
		}

		notifications = append(
			notifications,
			map[string]interface{}{
				"id":      id,
				"userId":  userID,
				"channel": channel,
				"message": message,
				"status":  status,
			},
		)
	}

	return notifications, nil
}

func GetNotificationByID(id string) (map[string]interface{}, error) {

	var notificationID int64
	var userID string
	var channel string
	var message string
	var status string

	query := `
	SELECT id, user_id, channel, message, status
	FROM notifications
	WHERE id = $1
	`

	err := DB.QueryRow(query, id).Scan(
		&notificationID,
		&userID,
		&channel,
		&message,
		&status,
	)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":      notificationID,
		"userId":  userID,
		"channel": channel,
		"message": message,
		"status":  status,
	}, nil
}

func GetNotificationsByStatus(status string) ([]map[string]interface{}, error) {

	rows, err := DB.Query(`
		SELECT id, user_id, channel, message, status
		FROM notifications
		WHERE status = $1
	`, status)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notifications []map[string]interface{}

	for rows.Next() {

		var id int64
		var userID string
		var channel string
		var message string
		var notificationStatus string

		err := rows.Scan(
			&id,
			&userID,
			&channel,
			&message,
			&notificationStatus,
		)

		if err != nil {
			return nil, err
		}

		notifications = append(
			notifications,
			map[string]interface{}{
				"id":      id,
				"userId":  userID,
				"channel": channel,
				"message": message,
				"status":  notificationStatus,
			},
		)
	}

	return notifications, nil
}