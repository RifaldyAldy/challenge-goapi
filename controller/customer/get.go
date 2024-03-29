package customer

import (
	"database/sql"
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	db := config.DB
	response := Response{}
	data := Customer{}
	query := "SELECT id,name,phonenumber,address FROM mst_customer WHERE id = $1"

	err := db.QueryRow(query, c.Param("id")).Scan(&data.Id, &data.Name, &data.PhoneNumber, &data.Address)
	if err == sql.ErrNoRows {
		response.Message = "Error, tidak ada id " + c.Param("id")
		response.Data = err
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Success"
	response.Data = data
	c.JSON(http.StatusOK, response)

}
