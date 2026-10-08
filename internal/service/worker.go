package service

import "log"

func Worker(id int, jobs <-chan string) {
	for msg := range jobs {

		log.Printf("[WORKER %d] Processing: %s", id, msg)

		// Send SMS
		err := SendSMS("+917483935838", msg)
		if err != nil {
			log.Printf("[ERROR] SMS failed: %v", err)
		} else {
			log.Printf("[SUCCESS] SMS sent")
		}

		// Send Email
		err = SendEmail(msg)
		if err != nil {
			log.Printf("[ERROR] Email failed: %v", err)
		} else {
			log.Printf("[SUCCESS] Email sent")
		}
	}
}

func StartWorkerPool() chan string {
	jobs := make(chan string, 100)

	for i := 1; i <= 3; i++ {
		go Worker(i, jobs)
	}

	return jobs
}