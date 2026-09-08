package settings

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Set(orgID, key, value string) (*OrgSetting, error) {
	var s OrgSetting

	query := `
        INSERT INTO org_settings (org_id, key, value)
        VALUES ($1, $2, $3)
        ON CONFLICT (org_id, key)
        DO UPDATE SET value = EXCLUDED.value
        RETURNING id, org_id, key, value
    `

	err := r.db.QueryRow(query, orgID, key, value).Scan(
		&s.ID, &s.OrgID, &s.Key, &s.Value,
	)

	return &s, err
}
