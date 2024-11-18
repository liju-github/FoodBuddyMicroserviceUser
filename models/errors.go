package model

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenGeneration = errors.New("failed to generate token")
	ErrInvalidToken    = errors.New("invalid token")
	ErrUserNotVerified = errors.New("user not verified")
	ErrInvalidCode     = errors.New("invalid verification code")
	ErrDuplicateEmail  = errors.New("email already exists")
)
