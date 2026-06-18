package spaces

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) CreateSpace(orgID, slug, name, description string, isPrivate bool) (*Space, error) {
	return s.repo.Create(orgID, slug, name, description, isPrivate)
}
