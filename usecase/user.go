package usecase

import (
	"architecture-test/domain/entity"
	"architecture-test/domain/repository"
)

func GetUsers(repo repository.UserRepository) entity.Users {
	users := repo.Index()
	// ドメインロジックやドメインバリデーションを書く
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	return users
}
