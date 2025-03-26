package service

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

//Мок для продюсур
func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.([]string), args.Error(1)
}

type MockPresenter struct {
	mock.Mock
}

//Мок для пресентер
func (m *MockPresenter) Present(data []string) error {
	args := m.Called()
	return args.Error(0)
}

//Проверка метода Run
func TestService_Run(t *testing.T) {
	mockProducer := new(MockProducer)
	mockPresenter := new(MockPresenter)
}

service := NewService(mockProducer, mockPresenter)
inputData := []string{"http://Some text"}
mockProducer.On("Produce").Return("http://Some text", nil)
mockPresenter.On("Present", []string{"http://Some text"}).Return(nil)

err := service.Run()
assert.Nil(t, err)

mockProducer.AssertExpectations(t)
mockPresenter.AssertExpectations(t)