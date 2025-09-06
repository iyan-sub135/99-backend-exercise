package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"public-api/models"
	"strconv"
	"time"
)

type UserService struct {
	baseURL string
	client  *http.Client
}

func NewUserService(baseURL string) *UserService {
	return &UserService{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetUserByID fetches a user by ID from user service
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	// Make request
	url := fmt.Sprintf("%s/users/%d", s.baseURL, id)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var response models.UserServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("user service error: %v", response.Errors)
	}

	return response.User, nil
}

// GetUsers fetches users from user service
func (s *UserService) GetUsers(pageNum, pageSize int) ([]models.User, error) {
	// Build query parameters
	params := url.Values{}
	params.Add("page_num", strconv.Itoa(pageNum))
	params.Add("page_size", strconv.Itoa(pageSize))

	// Make request
	url := fmt.Sprintf("%s/users?%s", s.baseURL, params.Encode())
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var response models.UserServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("user service error: %v", response.Errors)
	}

	return response.Users, nil
}

// CreateUser creates a user via user service
func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	// Prepare form data (user service expects form-encoded)
	data := url.Values{}
	data.Set("name", req.Name)

	// Make request
	url := fmt.Sprintf("%s/users", s.baseURL)
	resp, err := s.client.PostForm(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var response models.UserServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("user service error: %v", response.Errors)
	}

	return response.User, nil
}
