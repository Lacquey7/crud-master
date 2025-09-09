package services

import "inventory-app/internal/model"

func (s *Service) CreateMovies(m *model.Movie) error {
	if err := s.db.Create(&m).Error; err != nil {
		return err
	}
	return nil
}
