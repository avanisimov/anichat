package user

import "time"

type UserProfile struct {
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	AvatarURL  *string    `json:"avatarUrl,omitempty"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

type UsersMeData struct {
	ID             int64        `json:"id"`
	Email          string       `json:"email"`
	ProfileCreated bool         `json:"profileCreated"`
	Profile        *UserProfile `json:"profile,omitempty"`
}

type UpsertProfileRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}