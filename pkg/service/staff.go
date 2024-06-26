package service

import "ims-staff-api/pkg/repository"

type StaffService struct {
	repo repository.Staff
}

func NewStaffService(repo repository.Staff) *StaffService {
	return &StaffService{repo: repo}
}
