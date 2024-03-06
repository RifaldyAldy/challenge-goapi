package customer

import (
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	db := config.DB
	data := Customer{}
	response := Response{}
	c.ShouldBind(&data)
	data.Id = c.Param("id")

	err := ValidateHp(data.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error," + err.Error()})
		return
	}
	query := "UPDATE mst_customer SET name=$1, phonenumber=$2, address=$3 WHERE id = $4"

	_, err = db.Exec(query, data.Name, data.PhoneNumber, data.Address, data.Id)
	if err != nil {
		panic(err)
	}

	response.Message = "Success"
	response.Data = data
	c.JSON(http.StatusOK, response)

}
