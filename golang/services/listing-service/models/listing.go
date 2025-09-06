package models

type Listing struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint   `json:"user_id" gorm:"not null" binding:"required"`
	ListingType string `json:"listing_type" gorm:"not null" binding:"required,oneof=rent sale"`
	Price       int    `json:"price" gorm:"not null" binding:"required,min=1"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime:micro"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime:micro"`
}

func (Listing) TableName() string {
	return "listings"
}

type CreateListingRequest struct {
	UserID      uint   `form:"user_id" binding:"required"`
	ListingType string `form:"listing_type" binding:"required,oneof=rent sale"`
	Price       int    `form:"price" binding:"required,min=1"`
}

type GetListingsParams struct {
	PageNum  int   `form:"page_num,default=1" binding:"min=1"`
	PageSize int   `form:"page_size,default=10" binding:"min=1,max=100"`
	UserID   *uint `form:"user_id"`
}

type StandardResponse struct {
	Result bool     `json:"result"`
	Errors []string `json:"errors,omitempty"`
}

type GetListingsResponse struct {
	StandardResponse
	Listings []Listing `json:"listings"`
}

type CreateListingResponse struct {
	StandardResponse
	Listing *Listing `json:"listing,omitempty"`
}
