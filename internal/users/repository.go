package users

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(email, username, passwordHash string) (*User, error) {
	var u User

	query := `
        INSERT INTO users (email, username, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id, email, username, display_name, password_hash, created_at
    `

	err := r.db.QueryRow(query, email, username, passwordHash).Scan(
		&u.ID, &u.Email, &u.Username, &u.DisplayName,
		&u.PasswordHash, &u.CreatedAt,
	)

	return &u, err
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var u User

	query := `
        SELECT id, email, username, display_name, password_hash, created_at
        FROM users
        WHERE email = $1
    `

	err := r.db.QueryRow(query, email).Scan(
		&u.ID, &u.Email, &u.Username, &u.DisplayName,
		&u.PasswordHash, &u.CreatedAt,
	)

	return &u, err
}
