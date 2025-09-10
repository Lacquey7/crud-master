package services

import (
	"billing-app/internal/model"
)

func (s *Service) GetBillingId(id string) (*[]model.Orders, error) {
	var o []model.Orders

	if err := s.db.Where("user_id = ? ", id).Find(&o).Error; err != nil {
		return nil, err
	}

	return &o, nil
}
