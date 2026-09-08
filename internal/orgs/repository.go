package orgs

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(name, slug string) (*Org, error) {
	var org Org

	query := `
        INSERT INTO orgs (name, slug)
        VALUES ($1, $2)
        RETURNING id, slug, name, description, website_url, is_verified, created_at
    `

	err := r.db.QueryRow(query, name, slug).Scan(
		&org.ID, &org.Slug, &org.Name, &org.Description,
		&org.WebsiteURL, &org.IsVerified, &org.CreatedAt,
	)

	return &org, err
}
