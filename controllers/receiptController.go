/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package controllers

import (
	"fetch-rewards-receipt-processor-challenge/datastore"
	"fetch-rewards-receipt-processor-challenge/models"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	// Generate a random UUID (version 4)
	u, error := uuid.NewRandom()
	if error != nil {
		fmt.Println("Failed to generate UUID:", error)
		return
	}

	// Parse request body
	if err := c.BindJSON((&receipt)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid Input",
			"error": err,
		})
		return
	}

	// Validate request body
	validate := validator.New()
	err := validate.Struct(receipt)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(400, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	ds := datastore.Points
	ds[u.String()] = 100

	fmt.Println(ds)

	c.IndentedJSON(http.StatusCreated, receipt)

}
