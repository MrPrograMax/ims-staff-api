package service

import (
	"ims-staff-api/models"
	"ims-staff-api/pkg/repository"
)

type StaffService struct {
	repo repository.Staff
}

func NewStaffService(repo repository.Staff) *StaffService {
	return &StaffService{repo: repo}
}

func (s *StaffService) GetStaffList() ([]models.UserDto, error) {
	return s.repo.GetStaffList()
}

func (s *StaffService) GetWorkerTasksById(workerId int64) ([]models.Task, error) {
	return s.repo.GetWorkerTasksById(workerId)
}

func (s *StaffService) AssignTask() error {
	return s.repo.AssignTask()
}
