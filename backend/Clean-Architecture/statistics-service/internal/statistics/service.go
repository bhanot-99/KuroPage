package statistics

import (
	"strconv"

	"github.com/elliaaan/statistics-service/models"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func toStr(id uint64) string {
	return strconv.FormatUint(id, 10)
}

type UserStats struct {
	Total    uint32
	Orders   uint32
	Products uint32
}

func (s *Service) GetUserStatistics(userID uint64) (*UserStats, error) {
	var total, orders, products int64

	err := s.DB.
		Model(&models.Event{}).
		Where("data LIKE ?", "%\"user_id\":"+toStr(userID)+"%").
		Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = s.DB.
		Model(&models.Event{}).
		Where("type LIKE ?", "order.%").
		Where("data LIKE ?", "%\"user_id\":"+toStr(userID)+"%").
		Count(&orders).Error
	if err != nil {
		return nil, err
	}

	err = s.DB.
		Model(&models.Event{}).
		Where("type LIKE ?", "product.%").
		Where("data LIKE ?", "%\"user_id\":"+toStr(userID)+"%").
		Count(&products).Error
	if err != nil {
		return nil, err
	}

	return &UserStats{
		Total:    uint32(total),
		Orders:   uint32(orders),
		Products: uint32(products),
	}, nil
}
func (s *Service) GetOrderCountByUser(userID uint64) (uint32, error) {
	var count int64

	err := s.DB.
		Model(&models.Event{}).
		Where("type LIKE ?", "order.%").
		Where("data LIKE ?", "%\"user_id\":"+toStr(userID)+"%").
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint32(count), nil
}
