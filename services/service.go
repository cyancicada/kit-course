package services

import (
	"errors"
	"strings"
)

type UcenterServiceInterface interface {
	GetUser(string) (string, error)
	UserList() ([]string, error)
}

type UcenterService struct{}

func (UcenterService) GetUser(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}
func (UcenterService) UserList() ([]string, error) {

	return []string{"a", "b", "c"}, nil
}

var ErrEmpty = errors.New("empty input")

type ServiceMiddleware func(UcenterServiceInterface) UcenterServiceInterface
