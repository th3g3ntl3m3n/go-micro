package main

type EmailsDB map[string]string

type EmailRepository struct {
	DB EmailsDB
}

func New(DB EmailsDB) *EmailRepository {
	return &EmailRepository{DB}
}

type Repository interface {
	GetByUserID(id string) *Email
}

func (repo *EmailRepository) GetByUserID(id string) *Email {
	return repo.DB[id]
}
