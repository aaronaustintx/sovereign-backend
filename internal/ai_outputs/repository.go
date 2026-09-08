package ai_outputs

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(artifactID, model, prompt, output string) (*AIOutput, error) {
	var a AIOutput

	query := `
        INSERT INTO ai_outputs (artifact_id, model, prompt, output)
        VALUES ($1, $2, $3, $4)
        RETURNING id, artifact_id, model, prompt, output, created_at
    `

	err := r.db.QueryRow(query, artifactID, model, prompt, output).Scan(
		&a.ID, &a.ArtifactID, &a.Model, &a.Prompt, &a.Output, &a.CreatedAt,
	)

	return &a, err
}
