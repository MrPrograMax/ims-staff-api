package service

import (
	"ims-staff-api/models"
	"ims-staff-api/pkg/repository"
)

type StatusService struct {
	repo repository.Status
}

func (s *StatusService) GetStatusList() ([]models.TaskStatus, error) {
	return s.repo.GetStatusList()
}

func (s *StatusService) GetStatusIdByName(status string) (int64, error) {
	return s.repo.GetStatusIdByName(status)
}

func NewStatusService(repo repository.Status) *StatusService {
	return &StatusService{repo: repo}
}
