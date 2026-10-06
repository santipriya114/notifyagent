package handler

import (
	"log"
	"net/http"
	"sync"
)

var (
	requestCount = make(map[string]int)
	mutex        sync.Mutex
)

func RateLimit(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ip := r.Header.Get("X-Forwarded-For")

		if ip == "" {
			ip = "localhost"
		}

		mutex.Lock()
		requestCount[ip]++
		count := requestCount[ip]
		mutex.Unlock()

		log.Printf("[INFO] IP: %s Count: %d", ip, count)

		if count > 5 {

			log.Printf("[INFO] Rate limit exceeded for IP: %s", ip)

			http.Error(
				w,
				"Rate limit exceeded",
				http.StatusTooManyRequests,
			)

			return
		}

		next(w, r)
	}
}
