package ai_outputs

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) CreateOutput(artifactID, model, prompt, output string) (*AIOutput, error) {
	return s.repo.Create(artifactID, model, prompt, output)
}
