package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/andygodish/yamlres/backend/internal/service"
	"github.com/gorilla/mux"
)

var resumeService service.ResumeServiceInterface


// InitRoutes initializes the router with services
func InitRoutes(rs service.ResumeServiceInterface) *mux.Router {
    resumeService = rs
    return NewRouter()
}

// wrapResponseWriter is a wrapper for http.ResponseWriter to capture the status code
type wrapResponseWriter struct {
	http.ResponseWriter
	status int
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(jsonContentTypeMiddleware)

	r.HandleFunc("/healthz", healthCheckHandler).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/resume", getResumeHandler).Methods(http.MethodGet)

	return r
}

// healthCheckHandler responds to health check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

// getResumeHandler retrieves the resume data
func getResumeHandler(w http.ResponseWriter, r *http.Request) {
	resume, err := resumeService.GetResume()
	if err != nil {
		http.Error(w, `{"error":"failed_to_get_resume","message":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// Convert to JSON and send response
	jsonData, err := json.Marshal(resume)
	if err != nil {
		http.Error(w, `{"error":"json_encoding_failed","message":"Failed to encode resume"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// WriteHeader captures the status code before writing it
func (wrw *wrapResponseWriter) WriteHeader(code int) {
	wrw.status = code
	wrw.ResponseWriter.WriteHeader(code)
}

// newWrapResponseWriter creates a new wrapped response writer
func newWrapResponseWriter(w http.ResponseWriter) *wrapResponseWriter {
	return &wrapResponseWriter{w, http.StatusOK}
}

// loggingMiddleware logs information about each HTTP request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record start time
		start := time.Now()

		// Create a response writer wrapper to capture the status code
		wrw := newWrapResponseWriter(w)

		// Call the next handler
		next.ServeHTTP(wrw, r)

		// Calculate request duration
		duration := time.Since(start)

		// Log request details
		log.Printf(
			"METHOD: %s | PATH: %s | REMOTE: %s | STATUS: %d | DURATION: %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			wrw.status,
			duration,
		)
	})
}

// jsonContentTypeMiddleware sets the Content-Type header to application/json for all responses
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set content type header before calling the next handler
		w.Header().Set("Content-Type", "application/json")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
