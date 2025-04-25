package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andygodish/yamlres/backend/internal/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)
type MockResumeService struct {
	mock.Mock
}

func (m *MockResumeService) GetResume() (*service.Resume, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.Resume), args.Error(1)
}

func TestInitRoutes(t *testing.T) {
	mockService := new(MockResumeService)
	
	router := InitRoutes(mockService)
	
	assert.NotNil(t, router, "Router should not be nil")
	assert.Equal(t, mockService, resumeService, "resumeService should be set correctly")
}

func TestNewRouter(t *testing.T) {
	mockService := new(MockResumeService)
	resumeService = mockService // Set the global variable
	
	router := NewRouter()
	
	assert.NotNil(t, router, "Router should not be nil")
	
	routes := []struct {
		path   string
		method string
	}{
		{"/healthz", "GET"},
		{"/api/resume", "GET"},
	}
	
	for _, route := range routes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		match := &mux.RouteMatch{}
		assert.True(t, router.Match(req, match), "Route %s %s should be registered", route.method, route.path)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	require.NoError(t, err, "Failed to create request")
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)
	
	handler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")
	
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err, "Response should be valid JSON")
	
	assert.Equal(t, "ok", response["status"], "Response should have status 'ok'")
}

func TestGetResumeHandler_Success(t *testing.T) {
	mockService := new(MockResumeService)
	mockService.On("GetResume").Return(&service.Resume{
		Basics: service.Basics{
			Name:  "Test User",
			Email: "test@example.com",
		},
	}, nil)
	
	resumeService = mockService
	
	req, err := http.NewRequest("GET", "/api/resume", nil)
	require.NoError(t, err, "Failed to create request")
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getResumeHandler)
	
	handler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200 OK")
	
	var response service.Resume
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err, "Response should be valid JSON")
	
	assert.Equal(t, "Test User", response.Basics.Name, "Name should match")
	assert.Equal(t, "test@example.com", response.Basics.Email, "Email should match")
	
	mockService.AssertExpectations(t)
}

func TestGetResumeHandler_Error(t *testing.T) {
	mockService := new(MockResumeService)
	mockService.On("GetResume").Return(nil, errors.New("test error"))
	
	// Set the global service
	resumeService = mockService
	
	req, err := http.NewRequest("GET", "/api/resume", nil)
	require.NoError(t, err, "Failed to create request")
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getResumeHandler)
	
	handler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Status code should be 500")
	
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err, "Response should be valid JSON")
	
	assert.Equal(t, "failed_to_get_resume", response["error"], "Error code should match")
	assert.Contains(t, response["message"], "test error", "Error message should contain service error")
	
	mockService.AssertExpectations(t)
}

func TestLoggingMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	require.NoError(t, err, "Failed to create request")
	
	rr := httptest.NewRecorder()
	
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})
	
	middlewareHandler := loggingMiddleware(testHandler)
	
	middlewareHandler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")
	assert.Equal(t, "test", rr.Body.String(), "Response body should be unchanged")
}

func TestJsonContentTypeMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	require.NoError(t, err, "Failed to create request")
	
	rr := httptest.NewRecorder()
	
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})
	
	middlewareHandler := jsonContentTypeMiddleware(testHandler)
	
	middlewareHandler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type header should be set")
	assert.Equal(t, "test", rr.Body.String(), "Response body should be unchanged")
}

func TestWrapResponseWriter(t *testing.T) {
	originalWriter := httptest.NewRecorder()
	wrappedWriter := newWrapResponseWriter(originalWriter)
	
	wrappedWriter.WriteHeader(http.StatusNotFound)
	
	assert.Equal(t, http.StatusNotFound, wrappedWriter.status, "Status should be stored in wrapper")
	assert.Equal(t, http.StatusNotFound, originalWriter.Code, "Status should be passed to underlying writer")
}

