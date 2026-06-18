package events

import "time"

type Event struct {
	ID         string    `json:"id"`
	OrgID      string    `json:"org_id"`
	UserID     string    `json:"user_id"`
	ArtifactID string    `json:"artifact_id"`
	EventType  string    `json:"event_type"`
	Metadata   string    `json:"metadata"`
	CreatedAt  time.Time `json:"created_at"`
}
