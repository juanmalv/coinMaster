package webService

import (
	"github.com/juanmalv/coinMaster/api/src/api/constants"
	"log"
	"net/http"
	"time"

	"github.com/juanmalv/coinMaster/api/src/api/database"
	"github.com/juanmalv/coinMaster/api/src/api/models"

	"github.com/gin-gonic/gin"
)

var incomeCollection = database.Client.Database(constants.DBName).Collection(constants.IncomeOperation)

//NewIncome creates a new income id with details
func NewIncome(c *gin.Context) {
	var req models.Income

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request is invalid. Please check the data and try again"})
		return
	}

	err := req.ValidateIncomeRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	req.Date = time.Now().Format(time.RFC1123)

	insertResult, err := incomeCollection.InsertOne(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error while trying to create Income"})
		return
	}

	log.Println("Income saved with ID: ", insertResult.InsertedID)

	c.JSON(http.StatusOK, gin.H{"Income ID": insertResult.InsertedID, "message": "Income saved successfully!", "created_at": req.Date})
}

//UpdateIncome endpoint updates an existing income
func UpdateIncome(c *gin.Context) {

}
