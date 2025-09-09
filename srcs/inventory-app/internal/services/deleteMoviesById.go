package services

import "inventory-app/internal/model"

func (s *Service) DeleteMoviesById(id int) error {
	m := model.Movie{}

	if err := s.db.Delete(&m, id).Error; err != nil {
		return err
	}
	return nil
}
