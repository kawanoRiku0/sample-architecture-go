package infra

import "architecture-test/domain/entity"

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return new(UserRepositoryImpl)
}

func (repo *UserRepositoryImpl) Index() entity.Users {
	// ダミーデータを用意。本来はこの中でDBにクエリを投げたり、外部のAPIを叩いたりする。
	users := []*entity.User{
		{"1", "ichiro", 20},
		{"2", "ziro", 18},
	}
	return users
}

func (repo *UserRepositoryImpl) Get(id string) entity.User {
	return entity.User{}
}
