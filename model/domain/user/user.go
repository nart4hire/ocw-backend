package user

import "time"

type User struct {
	Email       string    `gorm:"primaryKey" json:"email"`
	Password    string    `json:"-"`
	Name        string    `json:"name"`
	Role        UserRole  `gorm:"type:user_role" json:"role"`
	IsActivated bool      `json:"activated"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
