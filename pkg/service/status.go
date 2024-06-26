package service

import (
	"ims-staff-api/models"
	"ims-staff-api/pkg/repository"
)

type StatusService struct {
	repo repository.Status
}

func NewStatusService(repo repository.Status) *StatusService {
	return &StatusService{repo: repo}
}

func (s StatusService) GetStatusList() ([]models.TaskStatus, error) {
	return s.repo.GetStatusList()
}
