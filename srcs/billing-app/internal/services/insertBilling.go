package services

import (
	"billing-app/internal/model"
	"encoding/json"
)

func (s *Service) InsertBilling(j []byte) error {

	var o model.Orders

	err := json.Unmarshal(j, &o)
	if err != nil {
		return err
	}

	if err := s.db.Create(&o).Error; err != nil {
		return err
	}

	return nil
}
