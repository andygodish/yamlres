package service

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewResumeService(t *testing.T) {
	expectPath := "test-path.yaml"

	service := NewResumeService(expectPath)

	assert.NotNil(t, service, "service should not be nil")
	assert.Equal(t, expectPath, service.filePath, "file path should be set correctly")
}

func TestGetResume_FileNotFound(t *testing.T) {
	service := NewResumeService("non-existent-file.yaml")

	resume, err := service.GetResume()

	assert.Nil(t, resume, "Resume should be nil when file not found")
	assert.Error(t, err, "Error should be returned when file not found")
	assert.Contains(t, err.Error(), "resume file not found", "Error message should indicate file not found")
}

func TestGetResume_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "valid_*.yaml")
	require.NoError(t, err, "Failed to create temp file")
	defer os.Remove(tmpFile.Name())

	validYAML := `
basics:
  name: John Doe
  email: john@example.com
work:
  - company: ACME Corp
    position: Developer
`
	_, err = tmpFile.Write([]byte(validYAML))
	require.NoError(t, err, "Failed to write to temp file")
	tmpFile.Close()

	service := NewResumeService(tmpFile.Name())

	resume, err := service.GetResume()

	assert.NoError(t, err, "Should not return an error for valid YAML file")
	assert.NotNil(t, resume, "Resume should not be nil")
	assert.Equal(t, "John Doe", resume.Basics.Name, "Name should be parsed correctly")
	assert.Equal(t, "john@example.com", resume.Basics.Email, "Email should be parsed correctly")
	assert.Len(t, resume.Work, 1, "Should have 1 work entry")
	assert.Equal(t, "ACME Corp", resume.Work[0].Company, "Company should be parsed correctly")
}

func TestGetResume_ReadFileError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping permission test on Windows")
	}

	tmpFile, err := os.CreateTemp("", "noperm_*.yaml")
	require.NoError(t, err, "Failed to create temp file")
	defer os.Remove(tmpFile.Name()) // Clean up after test

	// Write some content and close the file
	_, err = tmpFile.Write([]byte("some content"))
	require.NoError(t, err, "Failed to write to temp file")
	tmpFile.Close()

	// Remove read permissions
	err = os.Chmod(tmpFile.Name(), 0000)
	require.NoError(t, err, "Failed to change file permissions")

	service := NewResumeService(tmpFile.Name())

	// Act
	resume, err := service.GetResume()

	// Assert
	assert.Nil(t, resume, "Resume should be nil when file can't be read")
	assert.Error(t, err, "Error should be returned when file can't be read")
	assert.Contains(t, err.Error(), "failed to read resume file", "Error message should indicate read failure")

	// Restore permissions so the file can be deleted, see prior defer statement
	_ = os.Chmod(tmpFile.Name(), 0666)
}

func TestGetResume_InvalidYAML(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "invalid_*.yaml")
	require.NoError(t, err, "Failed to create temp file")
	defer os.Remove(tmpFile.Name()) // Clean up after test

	invalidYAML := `
basics:
  name: John Doe
  - this is invalid YAML
  email: john@example.com
`
	_, err = tmpFile.Write([]byte(invalidYAML))
	require.NoError(t, err, "Failed to write to temp file")
	tmpFile.Close()

	service := NewResumeService(tmpFile.Name())

	resume, err := service.GetResume()

	assert.Nil(t, resume, "Resume should be nil when YAML is invalid")
	assert.Error(t, err, "Error should be returned when YAML is invalid")
	assert.Contains(t, err.Error(), "failed to parse resume YAML", "Error message should indicate parsing failure")
}
