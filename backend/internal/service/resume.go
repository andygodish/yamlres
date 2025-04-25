// This should go in internal/service/resume.go
package service

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ResumeServiceInterface defines the methods a resume service must implement
type ResumeServiceInterface interface {
    GetResume() (*Resume, error)
}

// Resume represents the structure of a resume in JSON Resume schema format
type Resume struct {
	Basics       Basics        `yaml:"basics" json:"basics"`
	Work         []Work        `yaml:"work" json:"work"`
	Volunteer    []Volunteer   `yaml:"volunteer" json:"volunteer,omitempty"`
	Education    []Education   `yaml:"education" json:"education"`
	Skills       []Skill       `yaml:"skills" json:"skills"`
	Certificates []Certificate `yaml:"certificates" json:"certificates,omitempty"`
	// Add other fields as needed
}

// Basics contains basic information
type Basics struct {
	Name  string `yaml:"name" json:"name"`
	Email string `yaml:"email" json:"email,omitempty"`
	// Add other fields as needed
}

// Work experience
type Work struct {
	Company    string   `yaml:"company" json:"company"`
	Position   string   `yaml:"position" json:"position"`
	StartDate  string   `yaml:"startDate" json:"startDate,omitempty"`
	EndDate    string   `yaml:"endDate" json:"endDate,omitempty"`
	Highlights []string `yaml:"highlights" json:"highlights,omitempty"`
}

// Volunteer experience
type Volunteer struct {
	Organization string `yaml:"organization" json:"organization"`
	StartDate    string `yaml:"startDate" json:"startDate,omitempty"`
	EndDate      string `yaml:"endDate" json:"endDate,omitempty"`
}

// Education background
type Education struct {
	Institution string `yaml:"institution" json:"institution"`
	StudyType   string `yaml:"studyType" json:"studyType,omitempty"`
	StartDate   string `yaml:"startDate" json:"startDate,omitempty"`
	EndDate     string `yaml:"endDate" json:"endDate,omitempty"`
}

// Skill
type Skill struct {
	Name     string   `yaml:"name" json:"name"`
	Keywords []string `yaml:"keywords" json:"keywords,omitempty"`
}

// Certificate obtained
type Certificate struct {
	Name string `yaml:"name" json:"name"`
}

// ResumeService handles operations on resume data
type ResumeService struct {
	filePath string
}

// NewResumeService creates a new resume service
func NewResumeService(filePath string) *ResumeService {
	return &ResumeService{
		filePath: filePath,
	}
}

// GetResume retrieves the resume data from the YAML file
func (s *ResumeService) GetResume() (*Resume, error) {
	// Check if file exists
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("resume file not found: %s", s.filePath)
	}

	// Read the file
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read resume file: %v", err)
	}

	// Parse the YAML
	var resume Resume
	if err := yaml.Unmarshal(data, &resume); err != nil {
		return nil, fmt.Errorf("failed to parse resume YAML: %v", err)
	}

	return &resume, nil
}
