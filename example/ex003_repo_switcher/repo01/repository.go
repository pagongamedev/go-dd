package repo01

import (
	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/example/ex003_repo_switcher/service"
)

// NewRepository func
func NewRepository() (service.Repository, error) {
	r := repo{}
	return &r, nil
}

type repo struct {
}

// ============== Repo File ====================

// GetMessage func
func (repo *repo) GetMessage(str string) (*godd.Map, *godd.Error) {
	return &godd.Map{"Message": "Repo 01 : " + str}, nil
}
