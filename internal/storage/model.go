package storage

import "time"

type StorageLocation struct {
	ID          string    `json:"id"`
	ArtifactID  string    `json:"artifact_id"`
	StorageType string    `json:"storage_type"`
	URI         string    `json:"uri"`
	CreatedAt   time.Time `json:"created_at"`
}
