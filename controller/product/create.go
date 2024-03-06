package product

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := config.DB
	data := Product{}
	c.ShouldBind(&data)
	if err := ValidateUnit(&data.Unit); err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error", Data: map[string]any{"info": "Unit harus Kg/Buah"}})
		return
	}

	query := "INSERT INTO mst_product (name,price,unit) VALUES ($1,$2,$3) RETURNING id"

	err := db.QueryRow(query, data.Name, data.Price, data.Unit).Scan(&data.Id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, Response{Message: "Success", Data: data})
}

// validasi unit hanya kg dan buah
func ValidateUnit(s *string) error {
	*s = strings.ToLower(*s)
	if *s != "kg" && *s != "buah" {
		return errors.New("unit harus Kg/Buah")
	}
	fixUnit(s)
	return nil
}

// validasi untuk fix contoh inputan kg menjadi Kg saat ke database
func fixUnit(s *string) {
	temp := strings.Split(*s, "")
	temp[0] = strings.ToUpper(temp[0])
	*s = strings.Join(temp, "")
	fmt.Println(*s)
}
