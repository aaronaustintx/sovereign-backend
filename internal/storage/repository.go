package storage

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(artifactID, storageType, uri string) (*StorageLocation, error) {
	var s StorageLocation

	query := `
        INSERT INTO storage_locations (artifact_id, storage_type, uri)
        VALUES ($1, $2, $3)
        RETURNING id, artifact_id, storage_type, uri, created_at
    `

	err := r.db.QueryRow(query, artifactID, storageType, uri).Scan(
		&s.ID, &s.ArtifactID, &s.StorageType, &s.URI, &s.CreatedAt,
	)

	return &s, err
}
