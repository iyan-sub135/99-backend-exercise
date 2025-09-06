package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"not null" binding:"required"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:micro"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:micro"`
}

func (User) TableName() string {
	return "users"
}

type CreateUserRequest struct {
	Name string `form:"name" binding:"required"`
}

type GetUsersParams struct {
	PageNum  int `form:"page_num,default=1" binding:"min=1"`
	PageSize int `form:"page_size,default=10" binding:"min=1,max=100"`
}

type StandardResponse struct {
	Result bool     `json:"result"`
	Errors []string `json:"errors,omitempty"`
}

type GetUsersResponse struct {
	StandardResponse
	Users []User `json:"users"`
}

type GetUserResponse struct {
	StandardResponse
	User *User `json:"user,omitempty"`
}

type CreateUserResponse struct {
	StandardResponse
	User *User `json:"user,omitempty"`
}
