package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func (mw *Middleware) TimeDuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now().Unix()
		fmt.Println("start", start)

		next.ServeHTTP(w, r)
		end := time.Now().Unix()
		fmt.Println("end", end)

		fmt.Println("end-start", end-start)
	})
}
