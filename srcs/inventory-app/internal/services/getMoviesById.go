package services

import "inventory-app/internal/model"

func (s *Service) GetMoviesById(id int) (*model.Movie, error) {
	var m model.Movie

	if err := s.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return &m, nil
}
