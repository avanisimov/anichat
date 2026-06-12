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

type SessionUser struct {
	SessionID  int64
	UserID     int64
	Email      string
	Role       string
	CreatedAt  time.Time
	ExpiresAt  time.Time
}