package services

import (
	"errors"
	"inventory-app/internal/model"
)

func (s *Service) UpdateMovies(patch *model.Movie, current *model.Movie) error {
	changes := map[string]interface{}{}

	if patch.Title != "" {
		changes["title"] = patch.Title
	}
	if patch.Description != "" {
		changes["description"] = patch.Description
	}
	if len(changes) == 0 {
		return errors.New("no data to update")
	}

	return s.db.Model(&model.Movie{}).
		Where("id = ?", current.ID).
		Updates(changes).Error
}
