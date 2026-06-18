package space_memberships

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) AddMember(spaceID, userID, role string) (*SpaceMembership, error) {
	return s.repo.Add(spaceID, userID, role)
}
