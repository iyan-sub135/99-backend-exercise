package controllers

import (
	"net/http"
	"strconv"

	"listing-service/models"
	"listing-service/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListingController struct {
	repo *repositories.ListingRepository
}

func NewListingController(db *gorm.DB) *ListingController {
	return &ListingController{
		repo: repositories.NewListingRepository(db),
	}
}

// GetListings handles GET /listings
func (c *ListingController) GetListings(ctx *gin.Context) {
	var params models.GetListingsParams

	params.PageNum = 1
	params.PageSize = 10

	if pageNum := ctx.Query("page_num"); pageNum != "" {
		if val, err := strconv.Atoi(pageNum); err == nil && val > 0 {
			params.PageNum = val
		}
	}

	if pageSize := ctx.Query("page_size"); pageSize != "" {
		if val, err := strconv.Atoi(pageSize); err == nil && val > 0 && val <= 100 {
			params.PageSize = val
		}
	}

	if userID := ctx.Query("user_id"); userID != "" {
		if val, err := strconv.ParseUint(userID, 10, 32); err == nil {
			uid := uint(val)
			params.UserID = &uid
		}
	}

	listings, err := c.repo.GetListings(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.GetListingsResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{"Failed to retrieve listings"},
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GetListingsResponse{
		StandardResponse: models.StandardResponse{Result: true},
		Listings:         listings,
	})
}

// CreateListing handles POST /listings
func (c *ListingController) CreateListing(ctx *gin.Context) {
	var req models.CreateListingRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CreateListingResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{err.Error()},
			},
		})
		return
	}

	listing, err := c.repo.CreateListing(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.CreateListingResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{"Failed to create listing"},
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.CreateListingResponse{
		StandardResponse: models.StandardResponse{Result: true},
		Listing:          listing,
	})
}
