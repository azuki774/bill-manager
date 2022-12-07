package factory

import (
	"fmt"

	"azuki774/bill-manager/internal/mawinter"
	"azuki774/bill-manager/internal/repository"
)

func NewUsecaseMawinter(h *repository.HTTPClient, f *repository.FileLoader) (u *mawinter.UsecaseMawinter, err error) {
	l, err := NewLogger()
	if err != nil {
		fmt.Printf("failed to get logger: %v\n", err)
		return nil, err
	}

	return &mawinter.UsecaseMawinter{Logger: l, HTTPClient: h, FileLoader: f}, nil
}

func NewFileLoader() *repository.FileLoader {
	return &repository.FileLoader{}
}
