package infrastructure

import "github.com/stretchr/testify/mock"

type LoggerMock struct {
	mock.Mock
}

func (m LoggerMock) Error(s string) {
	m.Called(s)
	return
}
