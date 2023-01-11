package model

type BillAPIResponse struct {
	BillName string `json:"bill_name"`
	Price    int    `json:"price"`
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
