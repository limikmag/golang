package singleton

type singleton struct {
	count int
}

var instance *singleton

//GetInstance function it is a portal to get an instance of singleton struct
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
