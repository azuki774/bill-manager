# bill-manager API

## GET /
Health Check用

## GET /bill/{yyyymm}

### Response
```
[
    {
        category_name: "water",
        price: 1234,
    },
    {
        category_name: "elect",
        price: 1234,
    },
        {
        category_name: "gas",
        price: 123,
    },
]
```
