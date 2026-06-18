package settings

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) Set(orgID, key, value string) (*OrgSetting, error) {
	return s.repo.Set(orgID, key, value)
}
