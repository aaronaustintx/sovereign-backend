package artifacts

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(spaceID, createdBy, artifactType, title string) (*Artifact, error) {
	var a Artifact

	query := `
        INSERT INTO artifacts (space_id, created_by, type, title)
        VALUES ($1, $2, $3, $4)
        RETURNING id, space_id, created_by, type, title, created_at
    `

	err := r.db.QueryRow(query, spaceID, createdBy, artifactType, title).Scan(
		&a.ID, &a.SpaceID, &a.CreatedBy, &a.Type, &a.Title, &a.CreatedAt,
	)

	return &a, err
}
