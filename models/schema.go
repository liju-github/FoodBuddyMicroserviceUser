package model

type User struct {
	ID               string `gorm:"primaryKey;type:varchar(255)"`
	Email            string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHash     string `gorm:"type:varchar(255)"`
	Name             string `gorm:"type:varchar(255)"`
	PhoneNumber      uint64
	Reputation       int32
	VerificationCode string `gorm:"type:varchar(255)"`
	IsBanned         bool
	IsVerified       bool
}

type UserAddress struct {
	ID         string `gorm:"primaryKey;type:varchar(255)"`
	UserID     string `gorm:"type:varchar(255)"`
	StreetName string `gorm:"type:varchar(255)"`
	Locality   string `gorm:"type:varchar(255)"`
	State      string `gorm:"type:varchar(255)"`
	Pincode    string `gorm:"type:varchar(6)"`
}
