package service

import godd "github.com/pagongamedev/go-dd"

// Service interface
type Service interface {
	MessageRead(str string) (*godd.Map, *godd.Error)
}

// Repository interface
type Repository interface {
	GetMessage(str string) (*godd.Map, *godd.Error)
}

// ======== service.go ============

// NewService New Service
func NewService(repo Repository) (Service, error) {
	svc := DemoService{repo}
	return &svc, nil
}

// DemoService struct
type DemoService struct {
	repo Repository
}

// ============== API File ====================

//MessageRead func
func (svc *DemoService) MessageRead(str string) (*godd.Map, *godd.Error) {
	return svc.repo.GetMessage(str)
}
