package spaces

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(orgID, slug, name, description string, isPrivate bool) (*Space, error) {
	var s Space

	query := `
        INSERT INTO spaces (org_id, slug, name, description, is_private)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, org_id, slug, name, description, is_private, created_at
    `

	err := r.db.QueryRow(query, orgID, slug, name, description, isPrivate).Scan(
		&s.ID, &s.OrgID, &s.Slug, &s.Name, &s.Description,
		&s.IsPrivate, &s.CreatedAt,
	)

	return &s, err
}
