package spaces

import "time"

type Space struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPrivate   bool      `json:"is_private"`
	CreatedAt   time.Time `json:"created_at"`
}
