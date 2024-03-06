# Aplikasi Enigma Laundry

## Postman document

Untuk informasi tentang postman ada di url : https://documenter.getpostman.com/view/24168904/2sA2xe5a3A

## API Spec

### Customer API

#### Create Customer

Request :

- Method : `POST`
- Endpoint : `/customers`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "name": "Rifaldy",
    "phonenumber": "628123456789",
    "address": "JL Jasa raharja No.1 Simpang 3"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "message": "Success",
  "data": {
    "id": "1",
    "name": "Rifaldy",
    "phoneNumber": "628123456789",
    "address": "JL Jasa raharja No.1 Simpang 3"
}
}
```

#### Get Customer

Request :

- Method : GET
- Endpoint : `/customers/1`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
    "message": "Success",
    "data": {
        "id": "1",
        "name": "RIFALDY",
        "phoneNumber": "089612792131",
        "address": "JL DR Wahiddin"
    }
}
```

#### Update Customer

Request :

- Method : PUT
- Endpoint : `/customers/1`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "name": "Kenshi yonezhu",
    "phonenumber": "081234567890",
    "address": "JL Jonggol"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "message": "Success",
    "data": {
        "id": "1",
        "name": "Kenshi yonezhu",
        "phoneNumber": "081234567890",
        "address": "JL Jonggol"
    }
}
```

#### Delete Customer

Request :

- Method : DELETE
- Endpoint : `/customers/1`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "Success",
  "data": "OK"
}
```

### Product API

#### Create Product

Request :

- Method : POST
- Endpoint : `/products`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "name": "Laundry Bedcover",
    "price": 10000,
    "unit" : "kg"
}
```

Response :

- Status Code: 201 Created
- Body:

```json
{
    "message": "Success",
    "data": {
        "id": "1",
        "name": "Laundry Bedcover",
        "price": 10000,
        "unit": "Kg"
    }
}
```

#### List Product

Request : http://localhost:8080/products?productName=dry

- Method : GET
- Endpoint : `/products`
  - Header :
  - Accept : application/json
- Query Param :
  - productName : `dry`,

Response :

- Status Code : 200 OK
- Body:

```json
{
    "message": "Success",
    "data": [
        {
            "id": "1",
            "name": "Laundry + Setrika",
            "price": 20000,
            "unit": "Kg"
        },
        {
            "id": "2",
            "name": "Laundry jilbab",
            "price": 20000,
            "unit": "Kg"
        }
    ]
}
```

#### Product By Id

Request :

- Method : GET
- Endpoint : `/products/1`
- Header :
  - Accept : application/json

Response :

- Status Code: 200 OK
- Body :

```json
{
    "message": "Success",
    "data": {
        "id": "1",
        "name": "Laundry + Setrika",
        "price": 20000,
        "unit": "Kg"
    }
}
```

#### Update Product

Request :

- Method : PUT
- Endpoint : `/products/1`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "name": "Laundry Bedcover",
    "price": 20000,
    "unit" : "bUaH"
}
```

Response :

- Status Code: 200 OK
- Body :

```json
{
    "message": "Success",
    "data": {
        "id": "1",
        "name": "Laundry Bedcover",
        "price": 20000,
        "unit": "Buah"
    }
}
```

#### Delete Product

Request :

- Method : DELETE
- Endpoint : `/products/1`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
    "message": "Success",
    "data": "OK"
}
```

### Transaction API

#### Create Transaction

Request :

- Method : POST
- Endpoint : `/transactions`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
	"billDate": "06-02-2024",
	"entryDate": "06-02-2024",
	"finishDate": "10-02-2024",
	"employeeId": "1",
	"customerId": "1",
	"billDetails": [
		{
			"productId": "1",
			"qty": 10
		},
        {
            "productId": "2",
            "qty": 20
        }
	]
}
```

Request :

- Status Code: 201 Created
- Body :

```json
{
    "message": "Success",
    "data": {
        "id": "1",
        "billDate": "06-02-2024",
        "entryDate": "06-02-2024",
        "finishDate": "10-02-2024",
        "employeeId": "1",
        "customerId": "1",
        "billDetails": [
            {
                "productId": "1",
                "qty": 10
            },
            {
                "productId": "2",
                "qty": 20
            }
        ]
    }
}
```

#### Get Transaction

Request :

- Method : GET
- Endpoint : `/transactions/1`
- Header :
  - Accept : application/json
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "message": "SUCCESS",
    "data": {
        "id": "1",
        "billDate": "01-January-0001",
        "entryDate": "06-February-2024",
        "finishDate": "10-February-2024",
        "employee": {
            "id": "1",
            "name": "Rizky",
            "phoneNumber": "6287173717371",
            "address": "JL Jasa raharja No.1 Simpang 3"
        },
        "customer": {
            "id": "1",
            "name": "Kenshi yonezhu",
            "phoneNumber": "081234567890",
            "address": "JL Jonggol"
        },
        "billDetails": [
            {
                "id": "2",
                "billId": "1",
                "product": {
                    "id": "2",
                    "name": "Laundry jilbab",
                    "price": 20000,
                    "unit": "Kg"
                },
                "productPrice": 400000,
                "quantity": 20
            },
            {
                "id": "1",
                "billId": "1",
                "product": {
                    "id": "1",
                    "name": "Laundry Bedcover",
                    "price": 20000,
                    "unit": "Buah"
                },
                "productPrice": 200000,
                "quantity": 10
            }
        ],
        "totalPrice": 600000
    }
}
```

#### List Transaction

Pattern string date : `dd-MM-yyyy`

Request :

- Method : GET
- Endpoint : `/transactions`
- Header :
  - Accept : application/json
- Query Param :
  - startDate : string `optional`
  - endDate : string `optional`
  - productName : string `optional`
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "message": "SUCCESS",
    "data": [
        {
            "id": "1",
            "billDate": "01-January-0001",
            "entryDate": "06-February-2024",
            "finishDate": "10-February-2024",
            "employee": {
                "id": "1",
                "name": "Rizky",
                "phoneNumber": "6287173717371",
                "address": "JL Jasa raharja No.1 Simpang 3"
            },
            "customer": {
                "id": "1",
                "name": "Kenshi yonezhu",
                "phoneNumber": "081234567890",
                "address": "JL Jonggol"
            },
            "billDetails": [
                {
                    "id": "2",
                    "billId": "1",
                    "product": {
                        "id": "2",
                        "name": "Laundry jilbab",
                        "price": 20000,
                        "unit": "Kg"
                    },
                    "productPrice": 400000,
                    "quantity": 20
                },
                {
                    "id": "1",
                    "billId": "1",
                    "product": {
                        "id": "1",
                        "name": "Laundry Bedcover",
                        "price": 20000,
                        "unit": "Buah"
                    },
                    "productPrice": 200000,
                    "quantity": 10
                }
            ],
            "totalBill": 600000
        },
        {
            "id": "2",
            "billDate": "06-February-2024",
            "entryDate": "06-February-2024",
            "finishDate": "10-February-2024",
            "employee": {
                "id": "1",
                "name": "Rizky",
                "phoneNumber": "6287173717371",
                "address": "JL Jasa raharja No.1 Simpang 3"
            },
            "customer": {
                "id": "1",
                "name": "Kenshi yonezhu",
                "phoneNumber": "081234567890",
                "address": "JL Jonggol"
            },
            "billDetails": [
                {
                    "id": "4",
                    "billId": "2",
                    "product": {
                        "id": "2",
                        "name": "Laundry jilbab",
                        "price": 20000,
                        "unit": "Kg"
                    },
                    "productPrice": 400000,
                    "quantity": 20
                },
                {
                    "id": "3",
                    "billId": "2",
                    "product": {
                        "id": "1",
                        "name": "Laundry Bedcover",
                        "price": 20000,
                        "unit": "Buah"
                    },
                    "productPrice": 200000,
                    "quantity": 10
                }
            ],
            "totalBill": 600000
        }
    ]
}
```