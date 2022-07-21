package main

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UsersDB map[string]*User

type UserRepository struct {
	DB UsersDB
}

func New(DB UsersDB) *UserRepository {
	return &UserRepository{DB}
}

type Repository interface {
	GetByID(id string) *User
}

func (repo *UserRepository) GetByID(id string) *User {
	return repo.DB[id]
}
