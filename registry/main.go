package registry

import (
	"architecture-test/domain/repository"
	"architecture-test/infra"
)

type Registry interface {
	NewUserRepository() repository.UserRepository
}

type RegistryImpl struct {
}

func NewRegistryImpl() *RegistryImpl {
	return &RegistryImpl{}
}

func (re *RegistryImpl) NewUserRepository() repository.UserRepository {
	return infra.NewUserRepository()
}
