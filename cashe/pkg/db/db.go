package db

import (
	"context"
)

type Interface interface {
	SaveURL(ctx context.Context, url URL) error
	GetOriginal(ctx context.Context, s string) (string, error)
}

type URL struct {
	Short string `json:"short"`
	Orig  string `json:"orig"`
}
