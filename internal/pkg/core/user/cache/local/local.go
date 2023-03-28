package local

import (
	"sync"

	"github.com/pkg/errors"

	cachePkg "HW-1/internal/pkg/core/user/cache"
	"HW-1/internal/pkg/core/user/models"
)

var (
	ErrUserExists   = errors.New("user exists")
	ErrUserNotExist = errors.New("user does not exist")
)

func New() cachePkg.Interface {
	return &cache{
		data: map[string]models.User{},
	}
}

type cache struct {
	mu   sync.RWMutex
	data map[string]models.User
}

func (c *cache) List() []models.User {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]models.User, 0, len(c.data))
	for _, value := range c.data {
		result = append(result, value)
	}

	return result
}

func (c *cache) Add(u models.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[u.Name]; ok {
		return errors.Wrapf(ErrUserExists, "user-name: [%s]: ", u.Name)
	}
	c.data[u.Name] = u
	return nil
}

func (c *cache) Update(u models.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[u.Name]; !ok {
		return errors.Wrapf(ErrUserNotExist, "user-name: [%s]: ", u.Name)
	}
	c.data[u.Name] = u
	return nil
}

func (c *cache) Delete(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if _, ok := c.data[name]; !ok {
		return errors.Wrapf(ErrUserNotExist, "user-name: [%s]: ", name)
	}
	delete(c.data, name)
	return nil
}
