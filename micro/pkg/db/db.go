package db

import (
	"context"
)

type Interface interface {
	AddURL(ctx context.Context, url URL) error
	Redirect(ctx context.Context, s string) (string, error)
}

type URL struct {
	Short string `json:"short"`
	Orig  string `json:"orig"`
}
