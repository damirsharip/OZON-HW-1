package cache

import "HW-1/internal/pkg/core/user/models"

type Interface interface {
	Add(user models.User) error
	Delete(name string) error
	List() []models.User
	Update(user models.User) error
}
