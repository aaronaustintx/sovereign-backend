package artifacts

import "time"

type Artifact struct {
	ID        string    `json:"id"`
	SpaceID   string    `json:"space_id"`
	CreatedBy string    `json:"created_by"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
