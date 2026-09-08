package space_memberships

import "time"

type SpaceMembership struct {
	ID        string    `json:"id"`
	SpaceID   string    `json:"space_id"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
