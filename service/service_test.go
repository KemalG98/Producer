package service

import (
	"github.com/stretchr/testify/assert"
	"src/service/mocks"
	"testing"
)

// Проверка метода Run
func TestService_Run(t *testing.T) {
	mockProducer := new(mocks.Producer)
	mockPresenter := new(mocks.Presenter)
	service := NewService(mockProducer, mockPresenter)

	urls := []string{"http://Some text"}
	mockProducer.On("Produce").Return(urls, nil)
	mockPresenter.On("Present", urls).Return(nil)

	err := service.Run()
	assert.Nil(t, err)

	mockProducer.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}
func TestService_Mask(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Mask letters", "http://Some text", "http://****************"},
		{"Mask numbers", "http://123", "http://**********"},
		{"Mask empty string", "", ""},
		{"Mask *", "http://****", "http://***********"},
		{"Mask special characters", "http://&$!@", "http://***********"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.Mask(tt.input); got != tt.want {
				t.Errorf("Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
