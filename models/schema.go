package model

type User struct {
	ID               string
	Email            string
	PasswordHash     string
	Name             string
	PhoneNumber      uint64
	Reputation       int32
	VerificationCode string
	IsBanned         bool
	IsVerified       bool
}

type UserAddress struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	StreetName string
	Locality   string
	State      string
	Pincode    string
}
