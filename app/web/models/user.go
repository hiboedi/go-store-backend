package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	Name      string    `json:"name" gorm:"not null;type:varchar(50)"`
	Email     string    `json:"email" gorm:"not null;unique;type:varchar(100)"`
	Password  string    `json:"password" gorm:"not null;type:varchar(100)"`
	Stores    []Store   `json:"stores" `
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserCreate struct {
	Name     string `validate:"required,min=4,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=50"`
}

type UserUpdate struct {
	Name     string `validate:"required,min=4,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=50"`
}

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func ToUserReponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// func ToUserResponses(users []User) []UserResponse {
// 	var responses []UserResponse

// 	for _, user := range users {
// 		responses = append(responses, ToUserReponse(user))
// 	}
// 	return responses
// }
