package orgs

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		repo: NewRepository(db),
	}
}

func (s *Service) CreateOrg(name, slug string) (*Org, error) {
	return s.repo.Create(name, slug)
}
