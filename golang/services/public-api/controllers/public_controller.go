package controllers

import (
	"net/http"
	"strconv"

	"public-api/models"
	"public-api/services"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
	listingService *services.ListingService
	userService    *services.UserService
}

func NewPublicController(listingService *services.ListingService, userService *services.UserService) *PublicController {
	return &PublicController{
		listingService: listingService,
		userService:    userService,
	}
}

// GetListings handles GET /public-api/listings
// Returns listings enriched with user data
func (c *PublicController) GetListings(ctx *gin.Context) {
	// Parse query parameters
	pageNum := 1
	pageSize := 10
	var userID *uint

	if pageNumStr := ctx.Query("page_num"); pageNumStr != "" {
		if val, err := strconv.Atoi(pageNumStr); err == nil && val > 0 {
			pageNum = val
		}
	}

	if pageSizeStr := ctx.Query("page_size"); pageSizeStr != "" {
		if val, err := strconv.Atoi(pageSizeStr); err == nil && val > 0 && val <= 100 {
			pageSize = val
		}
	}

	if userIDStr := ctx.Query("user_id"); userIDStr != "" {
		if val, err := strconv.ParseUint(userIDStr, 10, 32); err == nil {
			uid := uint(val)
			userID = &uid
		}
	}

	// Get listings from listing service
	listings, err := c.listingService.GetListings(pageNum, pageSize, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": false,
			"error":  "Failed to retrieve listings",
		})
		return
	}

	// Enrich listings with user data
	enrichedListings := make([]models.EnrichedListing, 0, len(listings))

	// Collect unique user IDs
	userIDs := make(map[uint]bool)
	for _, listing := range listings {
		userIDs[listing.UserID] = true
	}

	// Fetch user data for all unique user IDs
	userCache := make(map[uint]models.User)
	for userID := range userIDs {
		user, err := c.userService.GetUserByID(userID)
		if err != nil {
			// If user not found, create a placeholder
			userCache[userID] = models.User{
				ID:   userID,
				Name: "Unknown User",
			}
		} else {
			userCache[userID] = *user
		}
	}

	// Build enriched listings
	for _, listing := range listings {
		enrichedListings = append(enrichedListings, models.EnrichedListing{
			ID:          listing.ID,
			ListingType: listing.ListingType,
			Price:       listing.Price,
			CreatedAt:   listing.CreatedAt,
			UpdatedAt:   listing.UpdatedAt,
			User:        userCache[listing.UserID],
		})
	}

	ctx.JSON(http.StatusOK, models.GetListingsResponse{
		Result:   true,
		Listings: enrichedListings,
	})
}

// CreateUser handles POST /public-api/users
func (c *PublicController) CreateUser(ctx *gin.Context) {
	var req models.CreateUserRequest

	// Bind JSON data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create user via user service
	user, err := c.userService.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.CreateUserResponse{
		User: *user,
	})
}

// CreateListing handles POST /public-api/listings
func (c *PublicController) CreateListing(ctx *gin.Context) {
	var req models.CreateListingRequest

	// Bind JSON data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create listing via listing service
	listing, err := c.listingService.CreateListing(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create listing",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.CreateListingResponse{
		Listing: *listing,
	})
}
