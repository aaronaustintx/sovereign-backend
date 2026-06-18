package events

import "database/sql"

type Service struct {
    repo *Repository
}

func NewService(db *sql.DB) *Service {
    return &Service{repo: NewRepository(db)}
}

func (s *Service) LogEvent(orgID, userID, artifactID, eventType, metadata string) (*Event, error) {
    return s.repo.Log(orgID, userID, artifactID, eventType, metadata)
}
