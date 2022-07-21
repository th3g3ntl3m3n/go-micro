package main

type EmailsDB map[string]string

type EmailRepository struct {
	DB EmailsDB
}

func NewEmailRepository(DB EmailsDB) *EmailRepository {
	return &EmailRepository{DB}
}

type Repository interface {
	GetByUserID(id string) string
}

func (repo *EmailRepository) GetByUserID(id string) string {
	return repo.DB[id]
}
