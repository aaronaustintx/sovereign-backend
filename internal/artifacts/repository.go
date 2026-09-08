package artifacts

import (
	"database/sql"
)

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

func (r *Repository) Get(id string) (*Artifact, error) {
	var a Artifact

	query := `
        SELECT id, space_id, created_by, type, title, created_at
        FROM artifacts
        WHERE id = $1
    `

	err := r.db.QueryRow(query, id).Scan(
		&a.ID, &a.SpaceID, &a.CreatedBy, &a.Type, &a.Title, &a.CreatedAt,
	)

	return &a, err
}

func (r *Repository) List(spaceID string) ([]Artifact, error) {
	query := `
        SELECT id, space_id, created_by, type, title, created_at
        FROM artifacts
        WHERE space_id = $1
        ORDER BY created_at DESC
    `

	rows, err := r.db.Query(query, spaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Artifact

	for rows.Next() {
		var a Artifact
		if err := rows.Scan(
			&a.ID, &a.SpaceID, &a.CreatedBy, &a.Type, &a.Title, &a.CreatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, a)
	}

	return list, nil
}

func (r *Repository) Feed(spaceID string) ([]Artifact, error) {
	// Same as List for now — later we merge AI outputs, comments, etc.
	return r.List(spaceID)
}
