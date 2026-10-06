package service

import "fmt"

func SendToDLQ(message string) {
	fmt.Println("DLQ MESSAGE:", message)
}