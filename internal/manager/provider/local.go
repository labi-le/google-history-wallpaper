package provider

import (
	"context"
	"time"
)

type Options struct {
	SearchPhrase string
	ByDate       time.Time
	ColorMode    ColorMode
}

type ColorMode uint8

const (
	Gloomy ColorMode = iota
	Normal
	Warm
	Cool
	Night
)

type Provider interface {
	Provide(ctx context.Context, opts Options) error
}

type LocalDir struct{}

func NewLocalDir() *LocalDir {
	return &LocalDir{}
}

func (l *LocalDir) Provide(ctx context.Context, opts Options) error {
	return nil
}
