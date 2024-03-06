package transaksidetailscreate

import (
	"database/sql"
	"net/http"
	"submission-project-enigma-laundry/config"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Transaction struct {
	Id          string        `json:"id"`
	BillDate    string        `json:"billDate"`
	EntryDate   string        `json:"entryDate"`
	FinishDate  string        `json:"finishDate"`
	EmployeeId  string        `json:"employeeId"`
	CustomerId  string        `json:"customerId"`
	BillDetails []BillDetails `json:"billDetails"`
}

type BillDetails struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"qty"`
}

func Create(c *gin.Context) {
	db := config.DB
	tx, _ := db.Begin()
	data := Transaction{}
	c.ShouldBind(&data)
	list := formatTanggalDB(data)
	//cek customer
	err := CheckCustomer(data.CustomerId, tx)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Message: "Error, Id Customer tidak ditemukan"})
		tx.Rollback()
		return
	}

	err = CheckEmployee(data.EmployeeId, tx)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Message: "Error, Id Employee tidak ditemukan"})
		tx.Rollback()
		return
	}

	queryTrans := "INSERT INTO transaction (billdate,entrydate,finishdate,employeeid,customerid) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	err = tx.QueryRow(queryTrans, list[0], list[1], list[2], data.EmployeeId, data.CustomerId).Scan(&data.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error, Rollback", Data: err})
		tx.Rollback()
		return
	}

	queryDetails := "INSERT INTO transaction_details (productid,quantity,billid) VALUES ($1,$2,$3)"
	for _, value := range data.BillDetails {
		err = checkProduct(value.ProductId, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Message: "Error, product id " + value.ProductId + " Tidak ada", Data: err})
			tx.Rollback()
			return
		}
		_, err := tx.Exec(queryDetails, value.ProductId, value.Quantity, data.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Message: "Error, Rollback", Data: err})
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, Response{Message: "Success", Data: data})
}

func CheckCustomer(id string, tx *sql.Tx) error {
	query := "SELECT FROM mst_customer WHERE id = $1"
	err := tx.QueryRow(query, id).Scan()
	if err == sql.ErrNoRows {
		return err
	}

	return nil
}

func CheckEmployee(id string, tx *sql.Tx) error {
	query := "SELECT FROM employee WHERE id = $1"
	err := tx.QueryRow(query, id).Scan()
	if err == sql.ErrNoRows {
		return err
	}

	return nil
}

func checkProduct(id string, tx *sql.Tx) error {
	query := "SELECT FROM mst_product WHERE id = $1"
	err := tx.QueryRow(query, id).Scan()
	if err == sql.ErrNoRows {
		return err
	}

	return nil
}

func formatTanggalDB(data Transaction) []string {
	list := []string{data.BillDate, data.EntryDate, data.FinishDate}
	for i, v := range list {
		tanggal, _ := time.Parse("02-01-2006", v)
		list[i] = tanggal.Format("2006-01-02")

	}
	return list
}
