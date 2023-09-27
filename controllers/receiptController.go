/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package controllers

import (
	"fetch-rewards-receipt-processor-challenge/datastore"

	"fetch-rewards-receipt-processor-challenge/models"

	"fetch-rewards-receipt-processor-challenge/helpers"

	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type resultItem struct {
	ID string `json:"id"`
}

func calculateRetailerPoints(name string) int {
	return helpers.CountAlphaNumeric(name)
}
func calculatePoints(receipt models.Receipt) int {
	res := calculateRetailerPoints(receipt.Retailer)
	return res
}

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

	points := calculatePoints(receipt)

	ds := datastore.Points
	ds[u.String()] = points

	fmt.Println(ds)
	result := resultItem{
		ID: u.String(),
	}

	c.IndentedJSON(http.StatusCreated, result)

}
