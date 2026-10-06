package repository

import "fmt"

func CreateTables() {

	query := `
	CREATE TABLE IF NOT EXISTS notifications (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(100),
		channel VARCHAR(50),
		message TEXT,
		status VARCHAR(20),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("notifications table ready")
}