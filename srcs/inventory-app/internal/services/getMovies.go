package services

import (
	"inventory-app/internal/model"
	"strings"
)

func (s *Service) GetMovies(title string) (*[]model.Movie, error) {
	println(title)
	var m []model.Movie

	if strings.TrimSpace(title) != "" {
		if err := s.db.Where("title = ?", title).Find(&m).Error; err != nil {
			return nil, err
		}
		return &m, nil
	}

	if err := s.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}
