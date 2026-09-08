package artifact_versions

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(artifactID string, version int, content string) (*ArtifactVersion, error) {
	var v ArtifactVersion

	query := `
        INSERT INTO artifact_versions (artifact_id, version_number, content)
        VALUES ($1, $2, $3)
        RETURNING id, artifact_id, version_number, content, created_at
    `

	err := r.db.QueryRow(query, artifactID, version, content).Scan(
		&v.ID, &v.ArtifactID, &v.VersionNumber, &v.Content, &v.CreatedAt,
	)

	return &v, err
}
