package space_memberships

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(spaceID, userID, role string) (*SpaceMembership, error) {
	var m SpaceMembership

	query := `
        INSERT INTO space_memberships (space_id, user_id, role)
        VALUES ($1, $2, $3)
        RETURNING id, space_id, user_id, role, created_at
    `

	err := r.db.QueryRow(query, spaceID, userID, role).Scan(
		&m.ID, &m.SpaceID, &m.UserID, &m.Role, &m.CreatedAt,
	)

	return &m, err
}
