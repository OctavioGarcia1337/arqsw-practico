package services

import "github.com/emikohmann/arq-software/ej-auth/domain"

type ServiceMock struct{}

func (ServiceMock) Login(cred domain.Credentials) (domain.Token, error) {
	//TODO implement me
	panic("implement me")
}
