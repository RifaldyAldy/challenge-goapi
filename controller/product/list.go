package product

import (
	"database/sql"
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	db := config.DB
	paramName := c.Query("productName")
	query := "SELECT id,name,price,unit FROM mst_product ORDER BY id ASC"
	var res *sql.Rows
	var err error
	if len(paramName) > 0 {
		query = "SELECT id,name,price,unit FROM mst_product WHERE name ILIKE '%' || $1 || '%' ORDER BY id ASC"
		res, err = db.Query(query, paramName)
	} else {
		res, err = db.Query(query)
	}
	datas := []Product{}
	if err != nil {
		panic(err)
	}

	for res.Next() {
		data := Product{}
		res.Scan(&data.Id, &data.Name, &data.Price, &data.Unit)
		datas = append(datas, data)
	}

	c.JSON(http.StatusOK, Response{Message: "Success", Data: datas})
}
