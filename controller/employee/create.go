package employee

import (
	"errors"
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db := config.DB
	data := Employee{}
	c.ShouldBind(&data)
	err := validateLenHp(data.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error," + err.Error(), Data: err})
		return
	}
	query := "INSERT INTO employee (name,phonenumber,address) VALUES ($1,$2,$3) RETURNING id"

	err = db.QueryRow(query, data.Name, data.PhoneNumber, data.Address).Scan(&data.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Error", Data: err})
		return
	}

	c.JSON(http.StatusBadRequest, Response{Message: "Success", Data: data})
}

func ValidateHp(hp string) error {
	kodeNegara := hp[0:3]
	no := hp[0:2]
	if no != "08" && kodeNegara != "628" {
		return errors.New(" PhoneNumber harus 08 atau 628")
	}
	err := validateLenHp(hp)
	if err != nil {
		return err
	}
	return nil
}

func validateLenHp(hp string) error {
	if len(hp) > 13 {
		return errors.New(" Nomor hp tidak boleh lebih dari 13")
	}
	return nil
}
