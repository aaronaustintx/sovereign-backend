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
func (r *Repository) Get(id string) (*Space, error) {
    var s Space

    query := `
        SELECT id, org_id, slug, name, description, is_private, created_at
        FROM spaces
        WHERE id = $1
    `

    err := r.db.QueryRow(query, id).Scan(
        &s.ID, &s.OrgID, &s.Slug, &s.Name,
        &s.Description, &s.IsPrivate, &s.CreatedAt,
    )

    return &s, err
}

func (r *Repository) List(orgID string) ([]Space, error) {
    query := `
        SELECT id, org_id, slug, name, description, is_private, created_at
        FROM spaces
        WHERE org_id = $1
        ORDER BY created_at DESC
    `

    rows, err := r.db.Query(query, orgID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var spaces []Space

    for rows.Next() {
        var s Space
        if err := rows.Scan(
            &s.ID, &s.OrgID, &s.Slug, &s.Name,
            &s.Description, &s.IsPrivate, &s.CreatedAt,
        ); err != nil {
            return nil, err
        }
        spaces = append(spaces, s)
    }

    return spaces, nil
}
