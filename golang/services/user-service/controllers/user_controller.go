package controllers

import (
	"net/http"
	"strconv"

	"user-service/models"
	"user-service/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	repo *repositories.UserRepository
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		repo: repositories.NewUserRepository(db),
	}
}

// GetUsers handles GET /users
func (c *UserController) GetUsers(ctx *gin.Context) {
	var params models.GetUsersParams

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

	users, err := c.repo.GetUsers(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.GetUsersResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{"Failed to retrieve users"},
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GetUsersResponse{
		StandardResponse: models.StandardResponse{Result: true},
		Users:            users,
	})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.GetUserResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{"Invalid user ID"},
			},
		})
		return
	}

	user, err := c.repo.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, models.GetUserResponse{
				StandardResponse: models.StandardResponse{
					Result: false,
					Errors: []string{"User not found"},
				},
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, models.GetUserResponse{
				StandardResponse: models.StandardResponse{
					Result: false,
					Errors: []string{"Failed to retrieve user"},
				},
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, models.GetUserResponse{
		StandardResponse: models.StandardResponse{Result: true},
		User:             user,
	})
}

// CreateUser handles POST /users
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req models.CreateUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.CreateUserResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{err.Error()},
			},
		})
		return
	}

	user, err := c.repo.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.CreateUserResponse{
			StandardResponse: models.StandardResponse{
				Result: false,
				Errors: []string{"Failed to create user"},
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.CreateUserResponse{
		StandardResponse: models.StandardResponse{Result: true},
		User:             user,
	})
}
