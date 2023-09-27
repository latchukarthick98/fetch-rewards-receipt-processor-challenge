/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package controllers

import (
	"fetch-rewards-receipt-processor-challenge/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

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

	c.IndentedJSON(http.StatusCreated, receipt)

}
