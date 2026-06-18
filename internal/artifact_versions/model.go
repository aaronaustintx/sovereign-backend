package artifact_versions

import "time"

type ArtifactVersion struct {
	ID            string    `json:"id"`
	ArtifactID    string    `json:"artifact_id"`
	VersionNumber int       `json:"version_number"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
}
