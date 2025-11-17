package di

import "github.com/lrs-rodrigues/estudo-golang/internal/service"

type Services struct {
	User *service.UserService
}

func NewServices(repos *Repositories) *Services {
	return &Services{
		User: service.NewUserService(repos.User),
	}
}
