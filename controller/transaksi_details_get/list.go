package transaksidetailsget

import (
	"net/http"
	"submission-project-enigma-laundry/config"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction_list struct {
	Message string         `json:"message"`
	Data    []Transactions `json:"data"`
}

type Transactions struct {
	Id          string         `json:"id"`
	BillDate    string         `json:"billDate"`
	EntryDate   string         `json:"entryDate"`
	FinishDate  string         `json:"finishDate"`
	Employee    Employee       `json:"employee"`
	Customer    Customer       `json:"customer"`
	BillDetails []Bill_details `json:"billDetails"`
	TotalBill   int            `json:"totalBill"`
}

func GetList(c *gin.Context) {
	db := config.DB
	datas := Transaction_list{}
	billQuery := "SELECT * FROM transaction"
	custQuery := "SELECT name,phonenumber,address FROM mst_customer WHERE id = $1"
	emplQuery := "SELECT name,phonenumber,address FROM employee WHERE id = $1"
	billDetails := "SELECT trans.id,trans.productid,trans.quantity,trans.billid,product.name,product.price,product.unit FROM transaction_details AS trans JOIN mst_product AS product ON trans.productid = product.id WHERE trans.billid = $1"
	dt, err := db.Query(billQuery)
	if err != nil {
		panic(err)
	}
	defer dt.Close()
	for dt.Next() {
		data := Transactions{}
		dt.Scan(&data.Id, &data.BillDate, &data.EntryDate, &data.FinishDate, &data.Employee.Id, &data.Customer.Id)
		formatTanggal(&data)
		_ = db.QueryRow(custQuery, data.Customer.Id).Scan(&data.Customer.Name, &data.Customer.PhoneNumber, &data.Customer.Address)
		_ = db.QueryRow(emplQuery, data.Employee.Id).Scan(&data.Employee.Name, &data.Employee.PhoneNumber, &data.Employee.Address)
		dt2, _ := db.Query(billDetails, data.Id)
		defer dt2.Close()
		for dt2.Next() {
			dataProduct := Bill_details{}
			dt2.Scan(&dataProduct.Id, &dataProduct.Product.Id, &dataProduct.Quantity, &dataProduct.BillId, &dataProduct.Product.Name, &dataProduct.Product.Price, &dataProduct.Product.Unit)
			dataProduct.ProductPrice = dataProduct.Product.Price * dataProduct.Quantity
			data.TotalBill += dataProduct.ProductPrice
			data.BillDetails = append(data.BillDetails, dataProduct)
		}
		datas.Data = append(datas.Data, data)
	}
	datas.Message = "SUCCESS"
	c.JSON(http.StatusOK, datas)
}

func formatTanggal(data *Transactions) {
	listTanggal := []*string{&data.BillDate, &data.EntryDate, &data.FinishDate}
	for _, v := range listTanggal {
		tanggal, _ := time.Parse("2006-01-02T15:04:05Z", *v)
		*v = tanggal.Format("02-January-2006")
	}

}
