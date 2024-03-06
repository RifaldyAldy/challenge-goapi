package customer

import (
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	db := config.DB
	response := Response{}
	query := "DELETE FROM mst_customer WHERE id = $1"

	_, err := db.Exec(query, c.Param("id"))
	if err != nil {
		panic(err)
	}

	response.Message = "Success"
	response.Data = "OK!"
	c.JSON(http.StatusOK, response)
}
