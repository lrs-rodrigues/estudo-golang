package di

import adapterhttp "github.com/lrs-rodrigues/estudo-golang/internal/adapter/http"

type Handler struct {
	User *adapterhttp.UserHandler
}

func NewHandler(services *Services) *Handler {
	return &Handler{
		User: adapterhttp.NewUserHandler(services.User),
	}
}
