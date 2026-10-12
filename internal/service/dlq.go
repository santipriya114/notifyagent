package service

import "log"

var DLQ = make(chan Message, 100)

func ProcessDLQ() {
	for msg := range DLQ {
		log.Printf("[DLQ] Failed Message: %s", msg.Content)
	}
}