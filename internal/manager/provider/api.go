package provider

import "context"

type ApiProvider interface {
	Provide(ctx context.Context, opts Options) error
}

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) Provide(ctx context.Context, opts Options) error {
	panic("implement me")
}
