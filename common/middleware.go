package common

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

// ExceptionHandlerMiddleware is a custom middleware to handle exceptions and map them to HTTP status codes.
func ExceptionHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				var statusCode int = http.StatusInternalServerError
				var errMessage string = "Internal Server Error"

				custErr, ok := r.(*CustomError)
				if ok {
					statusCode = custErr.StatusCode
					errMessage = custErr.Err.Error()
				}

				// Log the error
				log.Printf("Error: %v\nStack Trace:\n%s", r, debug.Stack())

				// Set the appropriate status code in the response
				w.WriteHeader(statusCode)

				// Write the error message in JSON format
				errorResponse := map[string]string{"message": errMessage}
				jsonResponse, err := json.Marshal(errorResponse)
				if err != nil {
					log.Printf("Failed to marshal error response: %v", err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonResponse)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
