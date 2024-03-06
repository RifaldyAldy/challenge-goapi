package product

import (
	"database/sql"
	"net/http"
	"strconv"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) {
	db := config.DB
	query := "SELECT id,name,price,unit FROM mst_product WHERE id = $1"
	data := Product{}
	if chekid := CheckId(c); chekid != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error, id harus berupa angka", Data: "Kosong"})
		return
	}
	err := db.QueryRow(query, c.Param("id")).Scan(&data.Id, &data.Name, &data.Price, &data.Unit)
	if err != nil {
		CheckIdIsExist(err, c)
		return
	}

	c.JSON(http.StatusOK, Response{Message: "Success", Data: data})
}

// validasi
func CheckId(c *gin.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return nil
}

func CheckIdIsExist(err error, c *gin.Context) {
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, Response{Message: "product id " + c.Param("id") + " tidak ada", Data: "404 Not Found"})
	}
}
