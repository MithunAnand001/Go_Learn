package service

import (
	"learn_service/dto"
	"learn_service/repository"
)

type TestingService interface {
	HealthCheck() (*dto.SuccessResponse,error)
}

type TestingServiceImpl struct {
	TestingRepo repository.TestingRepository
}

// NewTestingService creates a new instance of TestingServiceImpl.
func NewTestingService(testingRepo repository.TestingRepository) *TestingServiceImpl {
	return &TestingServiceImpl{TestingRepo: testingRepo}
}

// HealthCheck checks the health of the service.
func (s *TestingServiceImpl) HealthCheck() (*dto.SuccessResponse,error){
	response, err := s.TestingRepo.CacheCheck()
	// Implementation of health check logic goes here
	if err != nil{
		return nil,err
	}
	return response,nil
}