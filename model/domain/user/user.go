package user

import "time"

type User struct {
	Email       string `gorm:"primaryKey"`
	Password    string
	Name        string
	Role        UserRole `gorm:"type:user_role"`
	IsActivated bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (User) TableName() string {
	return "user"
}
