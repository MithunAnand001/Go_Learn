package repository

import "learn_service/dto"

type TestingStruct struct {
}

type TestingRepository interface {
	CacheCheck() (*dto.SuccessResponse, error)
}

func NewTestingRepository() TestingStruct {
	return TestingStruct{}
}

func (t TestingStruct) CacheCheck() (*dto.SuccessResponse,error){
	return nil,nil
}