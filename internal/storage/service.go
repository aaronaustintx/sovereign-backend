package storage

import "database/sql"

type Service struct {
	repo *Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) AddLocation(artifactID, storageType, uri string) (*StorageLocation, error) {
	return s.repo.Add(artifactID, storageType, uri)
}
