package service

import "log"

func Worker(
	id int,
	jobs chan Message,
	dlq chan Message,
) {
	for msg := range jobs {

		log.Printf("[WORKER %d] Processing: %s", id, msg.Content)

		// Send SMS
		err := SendSMS("+917483935838", msg.Content)
		if err != nil {
			log.Printf("[ERROR] SMS failed: %v", err)
		} else {
			log.Printf("[SUCCESS] SMS sent")
		}

		// Send Email
		err = SendEmail(msg.Content)
		if err != nil {
			log.Printf("[ERROR] Email failed: %v", err)
		} else {
			log.Printf("[SUCCESS] Email sent")
		}
	}
}

func StartWorkerPool() chan Message {

	jobs := make(chan Message, 100)

	for i := 1; i <= 3; i++ {
		go Worker(
			i,
			jobs,
			DLQ,
		)
	}

	go ProcessDLQ()

	return jobs
}