package factory

import (
	"fmt"

	"github.com/azuki774/bill-manager/internal/repository"
	"github.com/azuki774/bill-manager/internal/usecases"
)

func NewUsecaseMawinter(h *repository.HTTPClient, f *repository.FileLoader) (u *usecases.UsecaseMawinter, err error) {
	l, err := NewLogger()
	if err != nil {
		fmt.Printf("failed to get logger: %v\n", err)
		return nil, err
	}

	return &usecases.UsecaseMawinter{Logger: l, HTTPClient: h, FileLoader: f}, nil
}

func NewFileLoader() *repository.FileLoader {
	return &repository.FileLoader{}
}
