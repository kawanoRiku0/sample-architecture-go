package handler

import "architecture-test/domain/entity"

type Users struct {
	Users []*User `json:"users"`
}

func UsersDTO(users entity.Users) Users {
	var result Users
	for _, u := range users {
		result.Users = append(result.Users, UserDTO(u))
	}
	return result
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"email"`
}

func UserDTO(user *entity.User) *User {
	return &User{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}
}
