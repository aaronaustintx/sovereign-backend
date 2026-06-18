package artifact_versions

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) CreateVersion(artifactID string, version int, content string) (*ArtifactVersion, error) {
	return s.repo.Create(artifactID, version, content)
}
