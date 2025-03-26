package service

import (
	"testing"
)

type MockProducer interface {
	MockProduce() ([]string, error)
}


type MockPresenter interface{
	MockPresent([]string) error
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

func TestService_Mask(t *testing.T) {

	tests := []struct {
		name   string
		input string
		want   string
	}{
		{"Mask short text", "abc", "*"},
		{"Mask longer text", "Hello World", "*"},
		{"Mask empty string", "", ""},
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