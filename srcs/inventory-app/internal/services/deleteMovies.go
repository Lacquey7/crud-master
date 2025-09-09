package services

import "inventory-app/internal/model"

func (s *Service) DeleteMovies() error {
	var m []model.Movie

	if err := s.db.Find(&m).Error; err != nil {
		return err
	}

	s.db.Delete(&m)

	return nil
}
