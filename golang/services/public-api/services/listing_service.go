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

type ListingService struct {
	baseURL string
	client  *http.Client
}

func NewListingService(baseURL string) *ListingService {
	return &ListingService{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetListings fetches listings from listing service
func (s *ListingService) GetListings(pageNum, pageSize int, userID *uint) ([]models.Listing, error) {
	// Build query parameters
	params := url.Values{}
	params.Add("page_num", strconv.Itoa(pageNum))
	params.Add("page_size", strconv.Itoa(pageSize))
	if userID != nil {
		params.Add("user_id", strconv.FormatUint(uint64(*userID), 10))
	}

	// Make request
	url := fmt.Sprintf("%s/listings?%s", s.baseURL, params.Encode())
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call listing service: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var response models.ListingServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("listing service error: %v", response.Errors)
	}

	return response.Listings, nil
}

// CreateListing creates a listing via listing service
func (s *ListingService) CreateListing(req models.CreateListingRequest) (*models.Listing, error) {
	// Prepare form data (listing service expects form-encoded)
	data := url.Values{}
	data.Set("user_id", strconv.FormatUint(uint64(req.UserID), 10))
	data.Set("listing_type", req.ListingType)
	data.Set("price", strconv.Itoa(req.Price))

	// Make request
	url := fmt.Sprintf("%s/listings", s.baseURL)
	resp, err := s.client.PostForm(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to call listing service: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var response models.ListingServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("listing service error: %v", response.Errors)
	}

	return response.Listing, nil
}
