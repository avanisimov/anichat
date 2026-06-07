package auth

import "time"

type User struct {
	ID        int64
	Email     string
	Role      string
	CreatedAt time.Time
}

type UserSession struct {
	ID        		int64
	UserID    		int64
	RefreshToken    string
	ExpiresAt 		time.Time
}