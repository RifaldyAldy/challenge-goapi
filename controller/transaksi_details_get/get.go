package transaksidetailsget

import (
	"net/http"
	"submission-project-enigma-laundry/config"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Message string              `json:"message"`
	Data    Transaction_details `json:"data"`
}

type Transaction_details struct {
	Id          string         `json:"id"`
	BillDate    string         `json:"billDate"`
	EntryDate   string         `json:"entryDate"`
	FinishDate  string         `json:"finishDate"`
	Employee    Employee       `json:"employee"`
	Customer    Customer       `json:"customer"`
	BillDetails []Bill_details `json:"billDetails"`
	TotalPrice  int            `json:"totalPrice"`
}

type Bill_details struct {
	Id           string  `json:"id"`
	BillId       string  `json:"billId"`
	Product      Product `json:"product"`
	ProductPrice int     `json:"productPrice"`
	Quantity     int     `json:"quantity"`
}
type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
type Employee struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"unit"`
}

func GetTransaction(c *gin.Context) {
	var db = config.DB
	id := c.Param("id_bill")
	var employeeid, customerid int
	T_Query := "SELECT id,billdate,entrydate,finishdate,employeeid,customerid FROM transaction WHERE id = $1"
	DetailsProductQuery := "SELECT trans.id,trans.productid,trans.quantity,trans.billid,product.name,product.price,product.unit FROM transaction_details AS trans JOIN mst_product AS product ON trans.productid = product.id WHERE trans.billid = $1;"
	employeeQuery := "SELECT id,name,phonenumber,address FROM employee WHERE id = $1"
	customerQuery := "SELECT id,name,phonenumber,address FROM mst_customer WHERE id = $1"

	dataResult := Transaction{}

	// Scan transaction
	err := db.QueryRow(T_Query, id).Scan(&dataResult.Data.Id, &dataResult.Data.BillDate, &dataResult.Data.EntryDate, &dataResult.Data.FinishDate, &employeeid, &customerid)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]any{"message": "ERROR", "error": err})
	}
	formatTanggalGet(&dataResult)
	// Scan customer
	err = db.QueryRow(customerQuery, customerid).Scan(&dataResult.Data.Customer.Id, &dataResult.Data.Customer.Name, &dataResult.Data.Customer.PhoneNumber, &dataResult.Data.Customer.Address)
	if err != nil {
		panic(err)
	}
	err = db.QueryRow(employeeQuery, employeeid).Scan(&dataResult.Data.Employee.Id, &dataResult.Data.Employee.Name, &dataResult.Data.Employee.PhoneNumber, &dataResult.Data.Employee.Address)
	if err != nil {
		panic(err)
	}

	product, err := db.Query(DetailsProductQuery, dataResult.Data.Id)
	if err != nil {
		panic(err)
	}
	for product.Next() {
		dataTrans := Bill_details{}
		product.Scan(&dataTrans.Id, &dataTrans.Product.Id, &dataTrans.Quantity, &dataTrans.BillId, &dataTrans.Product.Name, &dataTrans.Product.Price, &dataTrans.Product.Unit)
		dataTrans.ProductPrice = dataTrans.Product.Price * dataTrans.Quantity
		dataResult.Data.BillDetails = append(dataResult.Data.BillDetails, dataTrans)
		dataResult.Data.TotalPrice += dataTrans.ProductPrice
	}
	dataResult.Message = "SUCCESS"
	c.JSON(200, dataResult)
}

func formatTanggalGet(data *Transaction) {
	listTanggal := []*string{&data.Data.BillDate, &data.Data.EntryDate, &data.Data.FinishDate}
	for _, v := range listTanggal {
		tanggal, _ := time.Parse("2006-01-02T15:04:05Z", *v)
		*v = tanggal.Format("02-January-2006")
	}

}
