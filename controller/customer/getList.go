package customer

import (
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	db := config.DB
	datas := []Customer{}
	query := "SELECT id,name,phonenumber,address FROM mst_customer"

	res, err := db.Query(query)
	for res.Next() {
		data := Customer{}
		res.Scan(&data.Id, &data.Name, &data.PhoneNumber, &data.Address)
		datas = append(datas, data)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Message: "Error", Data: err})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "Success", Data: datas})

}
