package pattern

import "sync"

type singleton struct {
	propertyA string
	propertyB int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func (s *singleton) SetPropertyA(propertyA string) {
	s.propertyA = propertyA
}

func (s *singleton) PropertyA() string {
	return s.propertyA
}

func (s *singleton) SetPropertyB(propertyB int) {
	s.propertyB = propertyB
}

func (s *singleton) PropertyB() int {
	return s.propertyB
}
