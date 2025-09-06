package models

// User represents user data from user service
type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Listing represents listing data from listing service
type Listing struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// EnrichedListing represents listing with user data for public API
type EnrichedListing struct {
	ID          uint   `json:"id"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	User        User   `json:"user"`
}

// Standard response structures for external APIs
type StandardResponse struct {
	Result bool     `json:"result,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// Public API request structures (JSON)
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateListingRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	ListingType string `json:"listing_type" binding:"required,oneof=rent sale"`
	Price       int    `json:"price" binding:"required,min=1"`
}

// Public API response structures
type GetListingsResponse struct {
	Result   bool              `json:"result"`
	Listings []EnrichedListing `json:"listings"`
}

type CreateUserResponse struct {
	User User `json:"user"`
}

type CreateListingResponse struct {
	Listing Listing `json:"listing"`
}

// Internal service response structures
type ListingServiceResponse struct {
	Result   bool      `json:"result"`
	Listings []Listing `json:"listings,omitempty"`
	Listing  *Listing  `json:"listing,omitempty"`
	Errors   []string  `json:"errors,omitempty"`
}

type UserServiceResponse struct {
	Result bool     `json:"result"`
	Users  []User   `json:"users,omitempty"`
	User   *User    `json:"user,omitempty"`
	Errors []string `json:"errors,omitempty"`
}
