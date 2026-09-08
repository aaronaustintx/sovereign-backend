package artifacts

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) CreateArtifact(spaceID, createdBy, artifactType, title string) (*Artifact, error) {
	return s.repo.Create(spaceID, createdBy, artifactType, title)
}

func (s *Service) GetArtifact(id string) (*Artifact, error) {
	return s.repo.Get(id)
}
func (s *Service) ListArtifacts(spaceID string) ([]Artifact, error) {
	return s.repo.List(spaceID)
}

func (s *Service) Feed(spaceID string) ([]Artifact, error) {
	return s.repo.Feed(spaceID)
}
