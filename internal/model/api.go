package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type BillAPIResponse struct {
	BillName string `json:"bill_name"`
	Price    int    `json:"price"`
}

func ValidYYYYMM(yyyymm string) (err error) {
	if err := validation.Validate(yyyymm, validation.Length(6, 6), is.Digit); err != nil {
		return fmt.Errorf("invalid YYYYMM: %w", ErrInvalidData)
	}
	return nil
}

func (b *BillElect) NewBillAPIResponse() BillAPIResponse {
	return BillAPIResponse{
		BillName: "elect",
		Price:    int(b.Price),
	}
}

func (b *BillGas) NewBillAPIResponse() BillAPIResponse {
	return BillAPIResponse{
		BillName: "gas",
		Price:    int(b.Price),
	}
}

func (b *BillWater) NewBillAPIResponse() BillAPIResponse {
	return BillAPIResponse{
		BillName: "water",
		Price:    int(b.Price),
	}
}
