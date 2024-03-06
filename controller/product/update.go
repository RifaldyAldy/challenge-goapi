package product

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var db *sql.DB = config.DB
	data := Product{}
	c.ShouldBind(&data)
	err := CheckParamId(c)
	data.Id = c.Param("id")
	query := "UPDATE mst_product SET name=$1, price=$2, unit=$3 WHERE id = $4"
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error, param product id harus angka", Data: "Kosong"})
		return
	}
	err = CheckData(data, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error, tidak ada data dengan id " + data.Id, Data: "Kosong"})
		return
	}
	if err = ValidateUnit(&data.Unit); err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error", Data: map[string]any{"info": "Unit harus Kg/Buah"}})
		return
	}
	_, err = db.Exec(query, data.Name, data.Price, data.Unit, data.Id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, Response{Message: "Success", Data: data})
}

// validasi apakah Param Id berupa angka
func CheckParamId(c *gin.Context) error {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	return nil
}

// validasi cek apakah data ada
func CheckData(data Product, db *sql.DB) error {
	query := "SELECT FROM mst_product WHERE id = $1"
	fmt.Println(data)
	err := db.QueryRow(query, data.Id).Scan()
	if err == sql.ErrNoRows {
		return err
	}

	return nil
}
