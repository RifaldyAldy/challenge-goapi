package product

import (
	"database/sql"
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var db = config.DB
	var id = c.Param("id")

	// validasi param harus angka
	err := CheckParamId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error, param id harus berupa angka saja"})
		return
	}
	query := "DELETE FROM mst_product WHERE id = $1 RETURNING id"
	err = db.QueryRow(query, id).Scan(&id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, Response{Message: "Error, tidak ada product id " + id})
		return
	}

	c.JSON(http.StatusNotFound, Response{Message: "Success", Data: "OK"})

}
