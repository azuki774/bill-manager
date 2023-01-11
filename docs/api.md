# bill-manager API

## GET /
Health Check用

## GET /bill/{yyyymm}

### Response
```
[
    {
        bill_name: "water",
        price: 1234,
    },
    {
        bill_name: "elect",
        price: 1234,
    },
        {
        bill_name: "gas",
        price: 123,
    },
]
```
