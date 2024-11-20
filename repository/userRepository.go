package repository

import (
	"errors"
	"fmt"
	"time"

	model "github.com/liju-github/FoodBuddyMicroserviceUser/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUserVerification(userID string, isVerified bool) error
	GetUserProfile(userID string) (*model.User, error)
	UpdateUser(user *model.User) error
	StoreVerificationCode(userID, code string) error
	GetVerificationCode(userID string) (string, error)
	CheckBan(userID string) (bool, error)
	UnBanUser(userID string) error
	BanUser(userID string) error
	GetAllUsers() ([]*model.User, error)

	AddAddress(userID string, address *model.UserAddress) (string, error)
	GetAddresses(userID string) ([]*model.UserAddress, error)
	EditAddress(userID, addressID string, address *model.UserAddress) error
	DeleteAddress(userID, addressID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Modify the existing userRepository to implement the new methods
func (r *userRepository) AddAddress(userID string, address *model.UserAddress) (string, error) {
	// Generate a unique address ID
	address.ID = fmt.Sprintf("addr_%d", time.Now().UnixNano())
	address.UserID = userID

	if err := r.db.Create(address).Error; err != nil {
		return "", fmt.Errorf("failed to add address: %w", err)
	}

	return address.ID, nil
}

func (r *userRepository) GetAddresses(userID string) ([]*model.UserAddress, error) {
	var addresses []*model.UserAddress
	if err := r.db.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve addresses: %w", err)
	}

	return addresses, nil
}

func (r *userRepository) EditAddress(userID, addressID string, address *model.UserAddress) error {
	// First, verify the address belongs to the user
	var existingAddress model.UserAddress
	if err := r.db.Where("id = ? AND user_id = ?", addressID, userID).First(&existingAddress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("address not found or does not belong to user")
		}
		return fmt.Errorf("failed to find address: %w", err)
	}

	// Update the address
	result := r.db.Model(&existingAddress).Updates(map[string]interface{}{
		"street_name": address.StreetName,
		"locality":    address.Locality,
		"state":       address.State,
		"pincode":     address.Pincode,
	})

	if result.Error != nil {
		return fmt.Errorf("failed to update address: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func (r *userRepository) DeleteAddress(userID, addressID string) error {
	// Delete the address, ensuring it belongs to the user
	result := r.db.Where("id = ? AND user_id = ?", addressID, userID).Delete(&model.UserAddress{})

	if result.Error != nil {
		return fmt.Errorf("failed to delete address: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("address not found or does not belong to user")
	}

	return nil
}

func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil
}

// CreateUser creates a new user record
func (r *userRepository) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByEmail retrieves a user by their email address
func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func (r *userRepository) GetUserByID(id string) (*model.User, error) {

	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &user, nil
}

// UpdateUserVerification updates the verification status of a user
func (r *userRepository) UpdateUserVerification(userID string, isVerified bool) error {
	result := r.db.Model(&model.User{}).Where("id = ?", userID).Update("is_verified", isVerified)
	if result.Error != nil {
		return fmt.Errorf("failed to update user verification: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// GetUserProfile retrieves the user profile by userID
func (r *userRepository) GetUserProfile(userID string) (*model.User, error) {
	var user model.User
	result := r.db.Model(&model.User{}).Where("id = ?", userID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, model.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user profile: %w", result.Error)
	}
	return &user, nil
}

// UpdateUser updates a user's information
func (r *userRepository) UpdateUser(user *model.User) error {
	result := r.db.Model(&model.User{}).Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":         user.Name,
			"phone_number": user.PhoneNumber,
		})
	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// StoreVerificationCode stores the verification code for a user
func (r *userRepository) StoreVerificationCode(userID, code string) error {
	result := r.db.Model(&model.User{}).Where("id = ?", userID).Update("verification_code", code)
	if result.Error != nil {
		return fmt.Errorf("failed to store verification code: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// GetVerificationCode retrieves the verification code for a user
func (r *userRepository) GetVerificationCode(userID string) (string, error) {
	var user model.User
	if err := r.db.Select("verification_code").Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", fmt.Errorf("failed to get verification code: %w", err)
	}
	return user.VerificationCode, nil
}

func (r *userRepository) CheckBan(userID string) (bool, error) {
	var user model.User
	if err := r.db.Select("is_banned").Where("id = ?", userID).First(&user).Error; err != nil {
		return true, errors.New("check ban failed")
	}
	if user.IsBanned {
		return true, nil
	}
	return false, nil
}

func (r *userRepository) BanUser(userID string) error {
	var user model.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	result := r.db.Model(&user).Update("is_banned", true)
	if result.Error != nil {
		return fmt.Errorf("failed to ban user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected, user may already be banned")
	}

	return nil
}

func (r *userRepository) UnBanUser(userID string) error {
	// Check if the user exists
	var user model.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return fmt.Errorf("failed to find user: %w", err)
	}

	if err := r.db.Model(&user).Updates(map[string]interface{}{"is_banned": false, "id": userID}).Error; err != nil {
		return fmt.Errorf("failed to unban user: %w", err)
	}

	return nil
}
