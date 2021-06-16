package services

import (
	"bot/models"
	"fmt"
)

//Category : ""
func (s *Service) Category(isParent string) ([]models.Category, error) {
	category, err := s.Daos.Category(isParent)
	if err != nil {
		fmt.Println(err.Error())
		return category, err
	}
	return category, nil
}
