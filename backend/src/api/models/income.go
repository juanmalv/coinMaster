package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanmalv/coinMaster/api/src/api/constants"
	"net/http"
)

//Income model details the data included on a basic income saving
type Income struct {
	Amount   *float64 `json:"amount",binding:"required"`
	Date     string   `json:"date"`
	Reason   string   `json:"reason"`
	Category string   `json:"category"`
}


func (operation Income)ValidateIncomeRequest(c *gin.Context) error {

	if operation.Amount == nil{
		c.JSON(http.StatusBadRequest, "Invalid Request")
		return fmt.Errorf("%v %v must be filled",constants.IncomeOperation,constants.AmountField)
	}

	if *operation.Amount <= 0 {
		c.JSON(http.StatusBadRequest, "Invalid Request")
		return fmt.Errorf("%v %v must be greater than 0",constants.IncomeOperation,constants.AmountField)
	}
	return nil
}