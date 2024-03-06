package employee

import (
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	db := config.DB
	datas := []Employee{}

	query := "SELECT id,name,phonenumber,address FROM employee ORDER BY id ASC"

	dt, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error", Data: err})
		return
	}
	for dt.Next() {
		data := Employee{}
		dt.Scan(&data.Id, &data.Name, &data.PhoneNumber, &data.Address)
		datas = append(datas, data)
	}

	c.JSON(http.StatusOK, Response{Message: "Success", Data: datas})
}
