package user

import (
	"github.com/pkg/errors"

	cachePkg "HW-1/internal/pkg/core/user/cache"
	localcachePkg "HW-1/internal/pkg/core/user/cache/local"

	"HW-1/internal/pkg/core/user/models"
)

var ErrValidation = errors.New("invalid data")

type Interface interface {
	Create(user models.User) error
	Update(user models.User) error
	Delete(name string) error
	List() []models.User
}

func New() Interface {
	return &core{
		cache: localcachePkg.New(),
	}
}

type core struct {
	cache cachePkg.Interface
}

func (c *core) Create(user models.User) error {
	if user.Name == "" {
		return errors.Wrap(ErrValidation, "field: [name] can not be empty")
	}
	if user.Password == "" {
		return errors.Wrap(ErrValidation, "field: [password] can not be empty")
	}

	//if err := c.cache.Add(user); err != nil {
	//	return errors.Wrap(err, "error while calling package: [cache], endpoint: [Add]")
	//}

	return c.cache.Add(user)
}

func (c *core) Update(user models.User) error {
	if user.Name == "" {
		return errors.Wrap(ErrValidation, "field: [name] can not be empty")
	}
	if user.Password == "" {
		return errors.Wrap(ErrValidation, "field: [password] can not be empty")
	}

	return c.cache.Update(user)
}

func (c *core) Delete(name string) error {
	if name == "" {
		return errors.Wrap(ErrValidation, "field: [name] can not be empty")
	}

	return c.cache.Delete(name)
}

func (c *core) List() []models.User {

	return c.cache.List()
}
