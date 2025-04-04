package service

type Services struct {
	EchoService IEchoService
}

func InitServices() *Services {
	return &Services{
		EchoService: newEchoService(),
	}
}
