package repository

import "architecture-test/domain/entity"

type UserRepository interface {
	Index() entity.Users
	Get(id string) entity.User
}
