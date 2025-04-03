package services

type Services struct {
	EchoService EchoService
}

func InitServices() *Services {
	return &Services{
		EchoService: NewEchoService(),
	}
}
