package events

import (
    "database/sql"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Log(orgID, userID, artifactID, eventType, metadata string) (*Event, error) {
    var e Event

    query := `
        INSERT INTO events (org_id, user_id, artifact_id, event_type, metadata)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, org_id, user_id, artifact_id, event_type, metadata, created_at
    `

    err := r.db.QueryRow(query, orgID, userID, artifactID, eventType, metadata).Scan(
        &e.ID,
        &e.OrgID,
        &e.UserID,
        &e.ArtifactID,
        &e.EventType,
        &e.Metadata,
        &e.CreatedAt,
    )

    return &e, err
}
