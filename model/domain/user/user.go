package user

type User struct {
	Username    string `gorm:"primaryKey"`
	Password    string
	Name        string
	Role        UserRole `gorm:"type:user_role"`
	IsActivated bool
	CreatedAt   int64 `gorm:"autoCreateTime:nano"`
	UpdatedAt   int64 `gorm:"autoUpdateTime:nano"`
}
