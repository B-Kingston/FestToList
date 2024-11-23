package main

import (
	"time"

	"gorm.io/gorm"
)

// SpotifyUser represents the structure of a Spotify user in your database.
type User struct {
	gorm.Model
	SpotifyUserID     *string    `gorm:"unique;not null"`    // Unique Spotify user ID
	DisplayName       *string    `gorm:"size:255"`           // User's display name
	Email             *string    `gorm:"size:255"`           // User's email address
	Country           *string    `gorm:"size:2"`             // ISO 3166-1 alpha-2 country code
	SubscriptionLevel *string    `gorm:"size:50"`            // Subscription level (e.g., free, premium)
	ProfileImageURL   *string    `gorm:"type:text"`          // URL to the user's profile image
	SpotifyURI        *string    `gorm:"size:255"`           // Spotify URI
	FollowersCount    *int       `gorm:"default:0"`          // Number of followers
	AccessToken       *string    `gorm:"type:text;not null"` // Access token for API calls
	RefreshToken      *string    `gorm:"type:text;not null"` // Refresh token for renewing access tokens
	TokenExpiry       *time.Time `gorm:"not null"`           // Expiry time of the access token
	Scopes            *string    `gorm:"type:text"`          // Scopes granted by the user (comma-separated)
	LastAuthenticated *time.Time `gorm:"autoUpdateTime"`     // Timestamp of last authentication
	// LastUpdated        *time.Time `gorm:"autoUpdateTime"`                   // Timestamp of last update in the database
	Queries []Query
}

type Query struct {
	gorm.Model
	QueryUpload   string
	QueryDateTime string
	UserID        uint
}
